package wol

import (
	"fmt"
	"net"
	"net/netip"

	"github.com/mdlayher/wol"
)

func Wake(addr string, target string) error {
	client, err := wol.NewClient()
	if err != nil {
		return err
	}

	macAddress, err := net.ParseMAC(target)
	if err != nil {
		return err
	}

	ipAddress := addr
	if netip.MustParseAddr(addr).IsPrivate() {
		ipAddress = "255.255.255.255"
	}

	return client.Wake(fmt.Sprintf("%s:9", ipAddress), macAddress)
}