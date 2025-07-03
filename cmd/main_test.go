package main

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestRetryCommand(t *testing.T) {
	origExiter := cli.OsExiter
	cli.OsExiter = func(int) {}
	defer func() { cli.OsExiter = origExiter }()

	app := setupApp()
	args := []string{"retry", "-c", "3", "-i", "1s", "-t", "1s", "--", "echo error 1>&2; false"}
	err := app.Run(args)
	if err == nil {
		t.Errorf("Expected error, got none")
	}

	args = []string{"retry", "-c", "3", "-v", "-i", "1s", "-t", "1s", "--", "echo error 1>&2; false"}
	err = app.Run(args)
	if err == nil {
		t.Errorf("Expected error, got none")
	}

	args = []string{"retry", "-c", "1", "-i", "1s", "-t", "1s", "--", "echo success; true"}
	err = app.Run(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

}
