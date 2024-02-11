package usecases

import (
	"backend/internal/app/api/computers/domain/entities"
	"backend/internal/app/api/computers/domain/repositories"
	"backend/internal/app/api/computers/interface/requests"
	"backend/internal/app/api/database"

	"gorm.io/gorm"
)

type ComputerUsecase interface {
	Create(requests.CreateComputerInput) (*entities.Computer, error)
	Update(uint, requests.UpdateComputerInput) (*entities.Computer, error)
	Delete(uint) (*entities.Computer, error)
	Get(uint) (*entities.Computer, error)
}

type computerUsecase struct {
	computerRepository repositories.ComputerRepository
}

func NewComputerUsecase(cr repositories.ComputerRepository) ComputerUsecase {
	return &computerUsecase{
		computerRepository: cr,
	}
}

func (cu computerUsecase) Create(input requests.CreateComputerInput) (*entities.Computer, error) {
	db := database.Get()
	computer := entities.Computer{
		HostName:   input.HostName,
		IPAddress:  input.IPAddress,
		MACAddress: input.MACAddress,
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		return cu.computerRepository.Create(tx, &computer)
	}); err != nil {
		return nil, err
	}

	return &computer, nil
}

func (cu computerUsecase) Update(id uint, input requests.UpdateComputerInput) (*entities.Computer, error) {
	var computer entities.Computer
	db := database.Get()

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := cu.computerRepository.FindOneById(tx, &computer, id); err != nil {
			return err
		}

		computer.HostName = input.HostName
		computer.IPAddress = input.IPAddress
		computer.MACAddress = input.MACAddress

		return cu.computerRepository.Update(tx, &computer)
	}); err != nil {
		return nil, err
	}

	return &computer, nil
}

func (cu computerUsecase) Delete(id uint) (*entities.Computer, error) {
	var computer entities.Computer
	db := database.Get()

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := cu.computerRepository.FindOneById(tx, &computer, id); err != nil {
			return err
		}

		return cu.computerRepository.Delete(tx, &computer)
	}); err != nil {
		return nil, err
	}

	return &computer, nil
}

func (cu computerUsecase) Get(id uint) (*entities.Computer, error) {
	var computer entities.Computer
	db := database.Get()

	if err := cu.computerRepository.FindOneById(db, &computer, id); err != nil {
		return nil, err
	}

	return &computer, nil
}
