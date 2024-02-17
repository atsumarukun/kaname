package usecases

import (
	"backend/internal/app/api/computers/domain/clients"
	"backend/internal/app/api/computers/domain/entities"
	"backend/internal/app/api/computers/domain/repositories"
	"backend/internal/app/api/computers/interface/requests"
	"backend/internal/app/api/database"
	"fmt"
	"net"
	"net/netip"

	"gorm.io/gorm"
)

type ComputerUsecase interface {
	Create(requests.CreateComputerInput) (*entities.Computer, error)
	Update(uint, requests.UpdateComputerInput) (*entities.Computer, error)
	Delete(uint) (*entities.Computer, error)
	Wake(uint) (*entities.Computer, error)
	Get(uint) (*entities.Computer, error)
	Search(*requests.SearchComputersQuery) (*[]entities.Computer, error)
}

type computerUsecase struct {
	computerClient     clients.ComputerClient
	computerRepository repositories.ComputerRepository
	db                 *gorm.DB
}

func NewComputerUsecase(cc clients.ComputerClient, cr repositories.ComputerRepository, db *gorm.DB) ComputerUsecase {
	return &computerUsecase{
		computerClient:     cc,
		computerRepository: cr,
		db:                 db,
	}
}

func (cu computerUsecase) Create(input requests.CreateComputerInput) (*entities.Computer, error) {
	computer := entities.Computer{
		HostName:   input.HostName,
		IPAddress:  input.IPAddress,
		MACAddress: input.MACAddress,
	}

	if err := cu.db.Transaction(func(tx *gorm.DB) error {
		return cu.computerRepository.Create(tx, &computer)
	}); err != nil {
		return nil, err
	}

	return &computer, nil
}

func (cu computerUsecase) Update(id uint, input requests.UpdateComputerInput) (*entities.Computer, error) {
	var computer entities.Computer

	if err := cu.db.Transaction(func(tx *gorm.DB) error {
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

	if err := cu.db.Transaction(func(tx *gorm.DB) error {
		if err := cu.computerRepository.FindOneById(tx, &computer, id); err != nil {
			return err
		}

		return cu.computerRepository.Delete(tx, &computer)
	}); err != nil {
		return nil, err
	}

	return &computer, nil
}

func (cu computerUsecase) Wake(id uint) (*entities.Computer, error) {
	var computer entities.Computer

	if err := cu.computerRepository.FindOneById(cu.db, &computer, id); err != nil {
		return nil, err
	}

	macAddress, err := net.ParseMAC(computer.MACAddress)
	if err != nil {
		return nil, err
	}

	ipAddress := computer.IPAddress
	if netip.MustParseAddr(computer.IPAddress).IsPrivate() {
		ipAddress = "255.255.255.255"
	}

	if err := cu.computerClient.Wake(fmt.Sprintf("%s:9", ipAddress), macAddress); err != nil {
		return nil, err
	}

	return &computer, nil
}

func (cu computerUsecase) Get(id uint) (*entities.Computer, error) {
	var computer entities.Computer

	if err := cu.computerRepository.FindOneById(cu.db, &computer, id); err != nil {
		return nil, err
	}

	return &computer, nil
}

func (cu computerUsecase) Search(query *requests.SearchComputersQuery) (*[]entities.Computer, error) {
	var computers []entities.Computer

	operator := func(s *string, d string) string {
		if s != nil {
			return *s
		} else {
			return d
		}
	}

	options := database.Options{
		Sort:  operator(query.Sort, "updated_at"),
		Order: operator(query.Order, "desc"),
	}

	if err := cu.computerRepository.Find(cu.db, &computers, &options); err != nil {
		return nil, err
	}

	return &computers, nil
}
