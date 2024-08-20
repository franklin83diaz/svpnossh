package pkg

import (
	"log"

	"github.com/vishvananda/netlink"
)

func CreateNetInterface() {

	// Create a network interface tun
	linkAttrs := netlink.NewLinkAttrs()
	linkAttrs.Name = "tun0"

	tun := &netlink.Tuntap{
		LinkAttrs: linkAttrs,
		Mode:      netlink.TUNTAP_MODE_TUN,
	}

	if err := netlink.LinkAdd(tun); err != nil {
		panic(err)
	}

	// Get the network interface tun
	link, err := netlink.LinkByName(tun.Name)
	if err != nil {
		log.Fatalf("Error getting the interface: %v", err)
	}

	// IP address
	addr, err := netlink.ParseAddr("10.20.30.1/30")
	if err != nil {
		log.Fatalf("Error parsing the ip addr: %v", err)
	}

	// Add the IP address to the network interface tun
	if err := netlink.AddrAdd(link, addr); err != nil {
		log.Fatalf("error adding the ip: %v", err)
	}

	// Activate the network interface tun
	if err := netlink.LinkSetUp(link); err != nil {
		log.Fatalf("Error activating network interface: %v", err)
	}

}
