package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SshConfgiOk() bool {

	// file sshd_config
	configFile := "/etc/ssh/sshd_config"

	// Open the file
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println("Error open file:", err)
		return false
	}
	defer file.Close()

	// Read the file
	buf := make([]byte, 1024)
	n, err := file.Read(buf)

	if err != nil {
		fmt.Println("Error read file:", err)
		return false
	}

	// Check if the file contains the string "PermitRootLogin yes"
	if string(buf[:n]) == "PermitTunnel yes" {
		return true
	}

	return false

}

func CheckSshdRunning() bool {

	cmd := exec.Command("systemctl", "is-active", "ssh")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error checking if the ssh service is running:", err)
		return false

	}

	if strings.TrimSpace(string(output)) == "active" {
		return true
	}

	return false

}
