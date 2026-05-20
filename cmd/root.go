// Package cmd provides the CLI commands for auto-git-commit-format.
/*
Copyright © 2026 Allan Golding Dwyre <allan.golding-dwyre@vidal.fr>
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const maxMessageLength = 30

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "agcf",
	Short: "A automatic git commit message formatter",
	Long:  `AGCF (Automatic Git Commit Format) is a CLI tool that helps developers create consistent and informative git commit messages by providing predefined templates for common commit types. With AGCF, you can easily generate commit messages that follow best practices, making it easier for your team to understand the purpose of each commit and maintain clean project history.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		makeCommitCmd("feat", "✨", "Nouvelle fonctionnalité"),
		makeCommitCmd("fix", "🐛", "Correction de bug"),
		makeCommitCmd("refactor", "♻️", "Refactoring"),
		makeCommitCmd("remove", "🔥", "Suppression de code"),
		makeCommitCmd("docs", "📝", "Documentation"),
		makeCommitCmd("build", "🔧", "Build / déploiement"),
		makeCommitCmd("deps", "⬆️", "Mise à jour dépendances"),
		makeCommitCmd("test", "✅", "Ajout de tests"),
	)
}

var execGitCommit = func(msg string) ([]byte, error) {
	return exec.Command("git", "commit", "-m", msg).CombinedOutput()
}

func makeCommitCmd(name, emoji, description string) *cobra.Command {
	return &cobra.Command{
		Use:   name + " [message]",
		Short: description,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			commit(cmd.OutOrStdout(), emoji, strings.Join(args, " "))
		},
	}
}

// ErrNoMessageProvided is returned when the commit message is empty.
var ErrNoMessageProvided = errors.New("no commit message provided")

// ErrMessageTooLong is returned when the commit message exceeds the maximum allowed length.
type ErrMessageTooLong struct{ CurrentLength int }

func (e *ErrMessageTooLong) Error() string {
	return fmt.Sprintf("message too long: %d characters (max %d)", e.CurrentLength, maxMessageLength)
}

// Is reports whether the target error is an ErrMessageTooLong with the same length.
func (e *ErrMessageTooLong) Is(target error) bool {
	_, ok := target.(*ErrMessageTooLong)
	return ok && e.CurrentLength == target.(*ErrMessageTooLong).CurrentLength
}

func validateMessage(message string) (string, error) {
	message = strings.TrimSpace(message)
	if message == "" {
		return "", ErrNoMessageProvided
	}
	if len(message) >= maxMessageLength {
		return "", &ErrMessageTooLong{CurrentLength: len(message)}
	}

	message = strings.ToUpper(message[:1]) + message[1:]

	return message, nil
}

func commit(w io.Writer, emoji, message string) {
	validatedMsg, err := validateMessage(message)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	full := emoji + " " + validatedMsg
	out, err := execGitCommit(full)
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		os.Exit(1)
	}
	_, err = fmt.Fprintln(w, string(out))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
