package repositories

import (
	"backend/internal/app/api/computers/domains/entities"
	"backend/internal/app/api/pkg/database"

	"gorm.io/gorm"
)

type ComputerRepository interface {
	Create(*gorm.DB, *entities.Computer) error
	Update(*gorm.DB, *entities.Computer) error
	Delete(*gorm.DB, *entities.Computer) error
	FindOneById(*gorm.DB, *entities.Computer, uint) error
	Find(*gorm.DB, *[]entities.Computer, *database.Options) error
}
