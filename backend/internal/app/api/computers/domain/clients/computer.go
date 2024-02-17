package clients

import "net"

type ComputerClient interface {
	Wake(string, net.HardwareAddr) error
}
