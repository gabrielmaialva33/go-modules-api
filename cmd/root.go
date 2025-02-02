package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-modules-api",
	Short: "A CLI for managing the Go Modules API",
	Long:  `This is a CLI application for managing the Go Modules API, including migrations and other administrative tasks.`,
}

// Execute runs the root command and any subcommands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newMigrateCmd())
	rootCmd.AddCommand(newSeedCmd())
}
