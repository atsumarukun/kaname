package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UpdateComputerInput struct {
	HostName   string `json:"host_name"`
	IPAddress  string `json:"ip_address"`
	MACAddress string `json:"mac_address"`
}

func (uci UpdateComputerInput) Validate() error {
	return validation.ValidateStruct(&uci,
		validation.Field(
			&uci.HostName,
			validation.Required.Error("Host name is required"),
		),
		validation.Field(
			&uci.IPAddress,
			validation.Required.Error("IP address is required"),
			is.IP.Error("Invalid IP address"),
		),
		validation.Field(
			&uci.MACAddress,
			validation.Required.Error("MAC address is required"),
			is.MAC.Error("Invalid MAC address"),
		),
	)
}
