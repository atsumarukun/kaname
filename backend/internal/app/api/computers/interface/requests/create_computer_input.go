package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CreateComputerInput struct {
	HostName   string `json:"host_name"`
	IPAddress  string `json:"ip_address"`
	MACAddress string `json:"mac_address"`
}

func (i CreateComputerInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.HostName,
			validation.Required.Error("Host name is required"),
		),
		validation.Field(
			&i.IPAddress,
			validation.Required.Error("IP address is required"),
			is.IP.Error("Invalid IP address"),
		),
		validation.Field(
			&i.MACAddress,
			validation.Required.Error("MAC address is required"),
			is.MAC.Error("Invalid MAC address"),
		),
	)
}
