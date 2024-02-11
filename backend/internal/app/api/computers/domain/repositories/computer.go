package repositories

import (
	"backend/internal/app/api/computers/domain/entities"

	"gorm.io/gorm"
)

type ComputerRepository interface {
	Create(*gorm.DB, *entities.Computer) error
	Update(*gorm.DB, *entities.Computer) error
	Delete(*gorm.DB, *entities.Computer) error
	FindOneById(*gorm.DB, *entities.Computer, uint) error
}
