package pkg

import (
	"fmt"
	"net"
)

func ValidateIp(ip string) bool {
	// Check if the IP address is valid
	if ip == "" {
		fmt.Println("Error: IP address is empty")
		return false
	}

	// Check if the IP address is valid
	if !net.ParseIP(ip).IsGlobalUnicast() && !net.ParseIP(ip).IsLoopback() {
		fmt.Println("Error: IP address is not valid")
		return false
	}

	return true
}
