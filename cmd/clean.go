package cmd

import (
	"fmt"
	"os"
	"svpnossh/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean the resources",
	Long:  `delete the network interface tun0 and netfilter rules.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.Cleanup()
		if err != nil {
			os.Exit(1)
		}
		fmt.Println("Resources cleaned up successfully")
		os.Exit(0)

	},
}
