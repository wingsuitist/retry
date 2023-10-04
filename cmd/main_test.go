package main

import (
	"testing"
)

func TestRetryCommand(t *testing.T) {
	app := setupApp()
	args := []string{"retry", "-c", "3", "-i", "1s", "-t", "1s", "--", "echo error 1>&2; false"}
	err := app.Run(args)
	if err == nil {
		t.Errorf("Expected error, got none")
	}

	args = []string{"retry", "-c", "1", "-i", "1s", "-t", "1s", "--", "echo success; true"}
	err = app.Run(args)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	args = []string{"retry", "-c", "3", "-v", "-i", "1s", "-t", "1s", "--", "echo error 1>&2; false"}
	err = app.Run(args)
	if err == nil {
		t.Errorf("Expected error, got none")
	}
}
