package responses

import (
	"backend/internal/app/api/computers/domain/entities"
	"time"
)

type Computer struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	HostName   string    `json:"host_name"`
	IPAddress  string    `json:"ip_address"`
	MACAddress string    `json:"mac_address"`
}

func FromEntity(entity *entities.Computer) *Computer {
	return &Computer{
		ID:         entity.ID,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
		HostName:   entity.HostName,
		IPAddress:  entity.IPAddress,
		MACAddress: entity.MACAddress,
	}
}
