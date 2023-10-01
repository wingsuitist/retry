package main

import (
	"testing"
	"time"
)

func TestRetryCommand(t *testing.T) {
	app := setupApp()
	args := []string{"", "-c", "3", "-i", "1s", "-t", "1s", "--", "false"}
	err := app.Run(args)
	if err == nil {
		t.Errorf("Expected error, got none")
	}

	args = []string{"", "-c", "1", "-i", "1s", "-t", "1s", "--", "true"}
	err = app.Run(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	args = []string{"", "-c", "3", "-i", "1s", "-t", "1s", "--", "sleep 2; true"}
	err = app.Run(args)
	if err == nil || err.Error() != "context deadline exceeded" {
		t.Errorf("Expected timeout error, got %v", err)
	}
}
