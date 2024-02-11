package persistences

import (
	"backend/internal/app/api/computers/domain/entities"
	"backend/internal/app/api/computers/domain/repositories"

	"gorm.io/gorm"
)

type computerPersistence struct{}

func NewComputerPersistence() repositories.ComputerRepository {
	return &computerPersistence{}
}

func (_ computerPersistence) Create(db *gorm.DB, computer *entities.Computer) error {
	return db.Create(computer).Error
}

func (_ computerPersistence) Update(db *gorm.DB, computer *entities.Computer) error {
	return db.Save(computer).Error
}

func (_ computerPersistence) FindOneById(db *gorm.DB, computer *entities.Computer, id uint) error {
	return db.First(&computer, id).Error
}
