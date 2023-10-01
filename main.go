package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
)

func setupApp() *cli.App {
	return &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "c",
				Value: 1,
				Usage: "Number of retries",
			},
			&cli.DurationFlag{
				Name:  "i",
				Value: 1 * time.Second,
				Usage: "Interval between retries",
			},
			&cli.DurationFlag{
				Name:  "t",
				Value: 1 * time.Second,
				Usage: "Timeout for each command run",
			},
			&cli.BoolFlag{
				Name:  "v",
				Value: false,
				Usage: "Verbose output",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("no command provided")
			}
			cmdStr := c.Args().Slice()[0]
			count := c.Int("c")
			interval := c.Duration("i")
			timeout := c.Duration("t")
			verbose := c.Bool("v")

			var lastOutput bytes.Buffer

			for i := 0; i < count; i++ {
				fmt.Printf("retrying %d of %d\n", i+1, count)
				ctx, cancel := context.WithTimeout(context.Background(), timeout)
				defer cancel()

				cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
				lastOutput.Reset()
				cmd.Stdout = &lastOutput
				cmd.Stderr = &lastOutput

				if verbose {
					cmd.Stdout = io.MultiWriter(os.Stdout, &lastOutput)
					cmd.Stderr = io.MultiWriter(os.Stderr, &lastOutput)
				}

				err := cmd.Run()
				if err == nil {
					return nil
				}

				time.Sleep(interval)
			}

			fmt.Println("Last command output:")
			fmt.Println(lastOutput.String())
			return fmt.Errorf("command failed after %d retries", count)
		},
	}
}

func main() {
	app := setupApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
