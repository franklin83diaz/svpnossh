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
		ok := pkg.SshConfgiOk()
		if !ok {
			//coloryelllow
			fmt.Print("\033[33m")
			//Warning
			fmt.Println("Warning: you need `PermitTunnel yes` in /etc/ssh/sshd_config for svpnossh working")
			//colorreset
			fmt.Print("\033[0m")
		}

		ok = pkg.CheckSshdRunning()
		if !ok {
			//coloryelllow
			fmt.Print("\033[33m")
			//Warning
			fmt.Println("Warning: you need the ssh service running for svpnossh working")
			//colorreset
			fmt.Print("\033[0m")
		}
		pkg.CreateNetInterface("10.20.30.1/30")
	},
}
