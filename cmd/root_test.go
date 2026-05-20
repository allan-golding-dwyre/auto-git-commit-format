package cmd //nolint:package-comments

import (
	"bytes"
	"errors"
	"strings"
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

func TestTypesIconsCommand(t *testing.T) {
	type typeToIcon struct {
		Type string
		Icon string
	}

	tests := []typeToIcon{
		{"feat", "✨"},
		{"fix", "🐛"},
		{"refactor", "♻️"},
		{"remove", "🔥"},
		{"docs", "📝"},
		{"build", "🔧"},
		{"test", "✅"},
		{"deps", "⬆️"},
	}

	originalExec := execGitCommit
	defer func() { execGitCommit = originalExec }()

	var capturedMsg string
	execGitCommit = func(msg string) ([]byte, error) {
		capturedMsg = msg
		return []byte(msg), nil
	}

	for _, test := range tests {
		t.Run(test.Type, func(t *testing.T) {
			capturedMsg = ""
			_, err := executeCommand(test.Type, "Test commit message")
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if !strings.Contains(capturedMsg, test.Icon) {
				t.Errorf("Expected commit message to contain icon %s, got %s", test.Icon, capturedMsg)
			}
		})
	}
}

func TestValidateMessage(t *testing.T) {
	t.Run("No message", func(t *testing.T) {
		_, err := validateMessage("")

		assertError(t, err, ErrNoMessageProvided)
	})

	t.Run("Only White Space", func(t *testing.T) {
		_, err := validateMessage("    ")

		assertError(t, err, ErrNoMessageProvided)
	})

	t.Run("Message too long", func(t *testing.T) {
		_, err := validateMessage(strings.Repeat("a", maxMessageLength+1))

		assertError(t, err, &ErrMessageTooLong{CurrentLength: maxMessageLength + 1})
	})

	t.Run("Capitalized Letter", func(t *testing.T) {
		msg, err := validateMessage("test message")
		want := "Test message"

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if msg != want {
			t.Errorf("Expected message to be %s, got '%s'", want, msg)
		}
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if !errors.Is(got, want) {
		t.Errorf("Expected error of type %q, got %q", want, got)
	}
}
