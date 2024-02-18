package temporaries

import (
	"backend/internal/app/api/computers/domains/clients"
	"net"

	"github.com/mdlayher/wol"
)

type computerTemporary struct {
	wolClient *wol.Client
}

func NewComputerTemporary(wolClient *wol.Client) clients.ComputerClient {
	return &computerTemporary{wolClient: wolClient}
}

func (ct computerTemporary) Wake(addr string, target net.HardwareAddr) error {
	return ct.wolClient.Wake(addr, target)
}
