package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "svpnossh",
	Short: "svpnossh is a CLI tool to connect to a VPN over SSH",
	Long:  `svpnoossh is a CLI tool to connect to a VPN over SSH.`,
	Run: func(cmd *cobra.Command, args []string) {
		// show help
		cmd.Help()
	},
}

// Execute is the entry point for the CLI
func Execute() error {
	return rootCmd.Execute()

}
