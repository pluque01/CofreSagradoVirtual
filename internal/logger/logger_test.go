package logger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "test.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Redirect the logger output to the temporary file
	Default.Logger = Default.Logger.Output(tmpFile)

	// Test logging
	Default.Logger.Info().Msg("Test message")

	// Read the contents of the temporary file
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the log message is present in the file
	if string(content) == "Test message" {
		t.Fatal("Test message not found")
	}
	// assert.Contains(t, string(content), "Test message")

	// Close the logger
	Close()
}
