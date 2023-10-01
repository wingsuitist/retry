package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
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
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("no command provided")
			}
			cmdStr := c.Args().Slice()[0]
			count := c.Int("c")
			interval := c.Duration("i")
			timeout := c.Duration("t")
			for i := 0; i < count; i++ {
				ctx, cancel := context.WithTimeout(context.Background(), timeout)
				defer cancel()
				cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
				err := cmd.Run()
				if err == nil {
					return nil
				}
				time.Sleep(interval)
			}
			return fmt.Errorf("command failed after %d retries", count)
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
