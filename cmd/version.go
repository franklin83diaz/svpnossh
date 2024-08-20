package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of svpnossh",
	Long: `Print the version number of svpnossh.
	Example:
	svpnossh version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("svpnossh v1.0.0")
	},
}
