package temporaries

import (
	"net"
	"testing"

	"github.com/mdlayher/wol"
)

func TestWake(t *testing.T) {
	IPAddress := "255.255.255.255:9"
	MACAddress, err := net.ParseMAC("00:00:00:00:00:00")
	if err != nil {
		t.Errorf(err.Error())
	}

	wolClient, err := wol.NewClient()
	if err != nil {
		t.Errorf(err.Error())
	}
	c := NewComputerTemporary(wolClient)

	if err := c.Wake(IPAddress, MACAddress); err != nil {
		t.Errorf(err.Error())
	}
}
