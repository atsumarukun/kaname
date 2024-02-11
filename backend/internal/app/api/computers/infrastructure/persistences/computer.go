package persistences

import (
	"backend/internal/app/api/computers/domain/entities"
	"backend/internal/app/api/computers/domain/repositories"

	"gorm.io/gorm"
)

type computerPersistence struct {}

func NewComputerPersistence() repositories.ComputerRepository {
	return &computerPersistence{}
}

func (_ computerPersistence) Create(db *gorm.DB, computer *entities.Computer) error {
	return db.Create(computer).Error
}
