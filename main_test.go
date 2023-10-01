package main

import (
	"testing"
)

func TestRetryCommand(t *testing.T) {
	app := setupApp()
	args := []string{"retry", "-c", "3", "-i", "1s", "-t", "1s", "--", "false"}
	err := app.Run(args)
	if err == nil {
		t.Errorf("Expected error, got none")
	}

	args = []string{"retry", "-c", "1", "-i", "1s", "-t", "1s", "--", "true"}
	err = app.Run(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// The timeout test would be difficult to implement here without mocking, skipping it for now.
}
