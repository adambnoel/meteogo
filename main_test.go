package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main
	main()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
