package cmd

import (
	"fmt"
	"svpnossh/pkg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)

}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "enable mode server for VPN with SSH",
	Long: `Enable mode server for VPN with SSH,
	create a network interface tuntap with name tun0 and up.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server mode")
		pkg.CreateNetInterface()

	},
}
