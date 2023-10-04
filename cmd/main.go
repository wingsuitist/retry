package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
	"strings"
	
	"github.com/urfave/cli/v2"
)

func setupApp() *cli.App {
	return &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "count",
				Aliases: []string{"c"},
				Value:   1,
				Usage:   "Number of retries",
			},
			&cli.DurationFlag{
				Name:    "interval",
				Aliases: []string{"i"},
				Value:   1 * time.Second,
				Usage:   "Interval between retries",
			},
			&cli.DurationFlag{
				Name:    "timeout",
				Aliases: []string{"t"},
				Value:   1 * time.Second,
				Usage:   "Timeout for each command run",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Value:   false,
				Usage:   "Verbose output",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("no command provided")
			}
			cmdStr := strings.Join(c.Args().Slice(), " ")
			count := c.Int("count")
			interval := c.Duration("interval")
			timeout := c.Duration("timeout")
			verbose := c.Bool("verbose")

			var err error
			var lastStdout, lastStderr bytes.Buffer

			for i := 0; i < count; i++ {
				if verbose {
					fmt.Fprintf(os.Stderr, "retrying %d of %d\n", i+1, count)
				}
				ctx, cancel := context.WithTimeout(context.Background(), timeout)
				defer cancel()

				cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
				cmd.Env = os.Environ()

				var stdout, stderr bytes.Buffer
				cmd.Stdout = &stdout
				cmd.Stderr = &stderr

				err = cmd.Run()

				lastStdout = stdout
				lastStderr = stderr

				if err == nil || verbose || i == count-1 {
					if lastStdout.Len() > 0 {
						os.Stdout.Write(lastStdout.Bytes())
					}
					if lastStderr.Len() > 0 {
						os.Stderr.Write(lastStderr.Bytes())
					}
				}

				if err == nil {
					break
				}
				time.Sleep(interval)
			}

			if err != nil {
				return fmt.Errorf("command failed after %d retries", count)
			}
			return nil
		},
	}
}

func main() {
	app := setupApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
