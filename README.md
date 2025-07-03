# The missing `retry` command

This is a Rust implementation of the `retry` utility. It allows running a command repeatedly until it succeeds or the retry count is exceeded.

```txt
USAGE:
   retry [OPTIONS] <COMMAND>...

ARGS:
   <COMMAND>...    Command to run

OPTIONS:
   -c, --count <COUNT>        Number of retries [default: 5]
   -i, --interval <INTERVAL>  Interval between retries [default: 3s]
   -t, --timeout <TIMEOUT>    Timeout for each command run [default: 1s]
   -v, --verbose              Verbose output
   -h, --help                 Print help
   -V, --version              Print version
```

## Building

Use [Cargo](https://doc.rust-lang.org/cargo/) to build the binary:

```sh
cargo build --release
```

The resulting executable will be available at `target/release/retry`.
