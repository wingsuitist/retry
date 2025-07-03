use clap::Parser;
use std::io::{self, Read, Write};
use std::process::{Command, ExitStatus, Stdio};
use std::time::Duration;
use wait_timeout::ChildExt;

#[derive(Parser, Debug)]
#[command(
    about = "The missing retry command",
    disable_help_subcommand = true,
    trailing_var_arg = true,
    version,
    author
)]
struct Args {
    /// Number of retries
    #[arg(short, long, default_value_t = 5)]
    count: u32,

    /// Interval between retries
    #[arg(short, long, value_parser = humantime::parse_duration, default_value = "3s")]
    interval: Duration,

    /// Timeout for each command run
    #[arg(short, long, value_parser = humantime::parse_duration, default_value = "1s")]
    timeout: Duration,

    /// Verbose output
    #[arg(short, long)]
    verbose: bool,

    /// Command to run
    #[arg(required = true)]
    command: Vec<String>,
}

fn run_once(cmd: &str, timeout: Duration) -> io::Result<(ExitStatus, Vec<u8>, Vec<u8>)> {
    let mut child = Command::new("bash")
        .arg("-c")
        .arg(cmd)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .spawn()?;

    let wait_res = child.wait_timeout(timeout)?;
    if wait_res.is_none() {
        let _ = child.kill();
    }
    let status = child.wait()?;

    let mut stdout = Vec::new();
    if let Some(mut out) = child.stdout.take() {
        let _ = out.read_to_end(&mut stdout);
    }
    let mut stderr = Vec::new();
    if let Some(mut err) = child.stderr.take() {
        let _ = err.read_to_end(&mut stderr);
    }

    Ok((status, stdout, stderr))
}

fn run(args: Args) -> Result<(), String> {
    let cmd_str = args.command.join(" ");
    let mut last_err = String::new();
    let mut last_stdout = Vec::new();
    let mut last_stderr = Vec::new();

    for i in 0..args.count {
        if args.verbose {
            eprintln!("retrying {} of {}", i + 1, args.count);
        }
        match run_once(&cmd_str, args.timeout) {
            Ok((status, out, err)) => {
                last_stdout = out;
                last_stderr = err;
                last_err = format!("{}", status);

                if status.success() || args.verbose || i == args.count - 1 {
                    if !last_stdout.is_empty() {
                        io::stdout().write_all(&last_stdout).unwrap();
                    }
                    if !last_stderr.is_empty() {
                        io::stderr().write_all(&last_stderr).unwrap();
                    }
                }

                if status.success() {
                    return Ok(());
                }
            }
            Err(e) => {
                last_err = e.to_string();
                if args.verbose || i == args.count - 1 {
                    if !last_stdout.is_empty() {
                        io::stdout().write_all(&last_stdout).unwrap();
                    }
                    if !last_stderr.is_empty() {
                        io::stderr().write_all(&last_stderr).unwrap();
                    }
                }
            }
        }
        std::thread::sleep(args.interval);
    }

    eprintln!("command failed after {} retries: {}", args.count, last_err);
    Err(last_err)
}

fn main() {
    let args = Args::parse();
    if run(args).is_err() {
        std::process::exit(1);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_retry_command() {
        let args = Args::parse_from([
            "retry",
            "-c",
            "3",
            "-i",
            "10ms",
            "-t",
            "10ms",
            "--",
            "bash",
            "-c",
            "echo error >&2; false",
        ]);
        assert!(run(args).is_err());

        let args = Args::parse_from([
            "retry",
            "-c",
            "3",
            "-v",
            "-i",
            "10ms",
            "-t",
            "10ms",
            "--",
            "bash",
            "-c",
            "echo error >&2; false",
        ]);
        assert!(run(args).is_err());

        let args = Args::parse_from([
            "retry",
            "-c",
            "1",
            "-i",
            "10ms",
            "-t",
            "10ms",
            "--",
            "bash",
            "-c",
            "echo success; true",
        ]);
        assert!(run(args).is_ok());
    }
}
