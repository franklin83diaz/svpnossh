package cmd

import (
	"fmt"
	"os/exec"
	"svpnossh/pkg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().String("host", "", "host to connect")
	clientCmd.Flags().String("user", "", "user to connect")

	clientCmd.MarkFlagRequired("host")
	clientCmd.MarkFlagRequired("user")

}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "enable mode client for VPN with SSH",
	Long: `Enable mode client for VPN with SSH,
	create a network interface tuntap with name tun0 and up.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("host")
		user, _ := cmd.Flags().GetString("user")

		ok := pkg.ValidateIp(ip)
		if !ok {
			panic("Error: IP address is not valid")
		}

		fmt.Println("client mode")
		pkg.CreateNetInterface("10.20.30.2/30")

		// Create the tunnel
		cmdSys := exec.Command("ssh", "-f", "-w", "0:0", user+"@"+ip, "true")
		err := cmdSys.Run()

		if err != nil {
			fmt.Println("Error creating the tunnel:", err)
		}
	},
}
