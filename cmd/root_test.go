package cmd

import (
	"testing"
)

func executeCommand(args ...string) (string, error) {
    buf := new(bytes.Buffer)
    rootCmd.SetOut(buf)
    rootCmd.SetErr(buf)
    rootCmd.SetArgs(args)
    err := rootCmd.Execute()
    return buf.String(), err
}

func TestRootCommand(t *testing.T) {
	output, err := executeCommand()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := "A automatic git commit message formatter"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}