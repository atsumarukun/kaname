package entities

import "gorm.io/gorm"

type Computer struct {
	gorm.Model
	HostName   string `gorm:"not null"`
	IPAddress  string `gorm:"size:15;not null"`
	MACAddress string `gorm:"size:17;not null"`
}
