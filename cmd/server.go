package cmd

import (
	"fmt"
	"log"
	"os/exec"

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

		// sudo ip tuntap add dev tun0 mode tun
		cmdSystem := exec.Command("sudo", "ip", "tuntap", "add", "dev", "tun0", "mode", "tun")
		if err := cmdSystem.Run(); err != nil {
			log.Fatal(err)
		}

	},
}
