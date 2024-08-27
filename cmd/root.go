package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "svpnossh",
	Short: "svpnossh is a CLI tool to connect to a VPN over SSH",
	Long:  `svpnoossh is a CLI tool to connect to a VPN over SSH.`,
	Run: func(cmd *cobra.Command, args []string) {
		// show help
		cmd.Help()
		os.Exit(0)
	},
}

// Execute is the entry point for the CLI
func Execute() error {
	return rootCmd.Execute()

}
