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
		HostName: input.HostName,
		IPAddress: input.IPAddress,
		MACAddress: input.MACAddress,
	}

	if err := db.Transaction(func (tx *gorm.DB) error {
		return cu.computerRepository.Create(tx, &computer)
	}); err != nil {
		return nil, err
	}

	return &computer, nil
}
