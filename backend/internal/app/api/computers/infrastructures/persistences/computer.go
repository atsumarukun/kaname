package persistences

import (
	"backend/internal/app/api/computers/domains/entities"
	"backend/internal/app/api/computers/domains/repositories"
	"backend/internal/app/api/database"
	"fmt"

	"gorm.io/gorm"
)

type computerPersistence struct{}

func NewComputerPersistence() repositories.ComputerRepository {
	return &computerPersistence{}
}

func (cp computerPersistence) Create(db *gorm.DB, computer *entities.Computer) error {
	return db.Create(computer).Error
}

func (cp computerPersistence) Update(db *gorm.DB, computer *entities.Computer) error {
	return db.Save(computer).Error
}

func (cp computerPersistence) Delete(db *gorm.DB, computer *entities.Computer) error {
	return db.Unscoped().Delete(computer).Error
}

func (cp computerPersistence) FindOneById(db *gorm.DB, computer *entities.Computer, id uint) error {
	return db.First(computer, id).Error
}

func (cp computerPersistence) Find(db *gorm.DB, computers *[]entities.Computer, options *database.Options) error {
	return db.Order(fmt.Sprintf("%s %s", options.Sort, options.Order)).Find(computers).Error
}
