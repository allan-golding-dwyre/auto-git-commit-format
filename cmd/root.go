/*
Copyright © 2026 Allan Golding Dwyre <allan.golding-dwyre@vidal.fr>

*/
package cmd

import (
	"os"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "agcf",
	Short: "A automatic git commit message formatter",
	Long: `AGCF (Automatic Git Commit Format) is a CLI tool that helps developers create consistent and informative git commit messages by providing predefined templates for common commit types. With AGCF, you can easily generate commit messages that follow best practices, making it easier for your team to understand the purpose of each commit and maintain clean project history.`,
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
        makeCommitCmd("feat",     "✨", "Nouvelle fonctionnalité"),
        makeCommitCmd("fix",      "🐛", "Correction de bug"),
        makeCommitCmd("refactor", "♻️", "Refactoring"),
        makeCommitCmd("remove",   "🔥", "Suppression de code"),
        makeCommitCmd("docs",     "📝", "Documentation"),
        makeCommitCmd("build",    "🔧", "Build / déploiement"),
        makeCommitCmd("deps",     "⬆️",  "Mise à jour dépendances"),
    )
}

func makeCommitCmd(name, emoji, description string) *cobra.Command {
    return &cobra.Command{
        Use:   name + " [message]",
        Short: description,
        Args:  cobra.MinimumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            commit(emoji, strings.Join(args, " "))
        },
    }
}

func commit(emoji, message string) {
    full := emoji + " " + message
    out, err := exec.Command("git", "commit", "-m", full).CombinedOutput()
    if err != nil {
        fmt.Fprintln(os.Stderr, string(out))
        os.Exit(1)
    }
    fmt.Println(string(out))
}
