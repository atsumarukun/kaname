package usecases

import (
	"backend/internal/app/api/computers/domains/entities"
	"backend/internal/app/api/computers/interfaces/requests"
	mock_clients "backend/internal/app/api/computers/mocks/clients"
	mock_repositories "backend/internal/app/api/computers/mocks/repositories"
	"backend/internal/app/api/pkg/database"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	input := requests.CreateComputerInput{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock_clients.NewMockComputerClient(ctrl)

	var err error
	r := mock_repositories.NewMockComputerRepository(ctrl)
	r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(err)

	u := NewComputerUsecase(c, r, database.Get())
	computer, err := u.Create(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, computer.HostName, input.HostName)
	assert.Equal(t, computer.IPAddress, input.IPAddress)
	assert.Equal(t, computer.MACAddress, input.MACAddress)
}

func TestUpdate(t *testing.T) {
	id := uint(1)
	input := requests.UpdateComputerInput{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock_clients.NewMockComputerClient(ctrl)

	var err error
	r := mock_repositories.NewMockComputerRepository(ctrl)
	r.EXPECT().FindOneById(gomock.Any(), gomock.Any(), id).Return(err)
	r.EXPECT().Update(gomock.Any(), gomock.Any()).Return(err)

	u := NewComputerUsecase(c, r, database.Get())
	computer, err := u.Update(id, input)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, computer.HostName, input.HostName)
	assert.Equal(t, computer.IPAddress, input.IPAddress)
	assert.Equal(t, computer.MACAddress, input.MACAddress)
}

func TestDelete(t *testing.T) {
	id := uint(1)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock_clients.NewMockComputerClient(ctrl)

	var err error
	r := mock_repositories.NewMockComputerRepository(ctrl)
	r.EXPECT().FindOneById(gomock.Any(), gomock.Any(), id).Return(err)
	r.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(err)

	u := NewComputerUsecase(c, r, database.Get())
	_, err = u.Delete(id)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestWake(t *testing.T) {
	id := uint(1)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	c := mock_clients.NewMockComputerClient(ctrl)
	c.EXPECT().Wake(gomock.Any(), gomock.Any()).Return(err)

	r := mock_repositories.NewMockComputerRepository(ctrl)
	r.EXPECT().FindOneById(gomock.Any(), gomock.Any(), id).Do(func(db *gorm.DB, computer *entities.Computer, id uint) {
		computer.IPAddress = "0.0.0.0"
		computer.MACAddress = "00:00:00:00:00:00"
	}).Return(err)

	u := NewComputerUsecase(c, r, database.Get())
	_, err = u.Wake(id)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	id := uint(1)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock_clients.NewMockComputerClient(ctrl)

	var err error
	r := mock_repositories.NewMockComputerRepository(ctrl)
	r.EXPECT().FindOneById(gomock.Any(), gomock.Any(), id).Return(err)

	u := NewComputerUsecase(c, r, database.Get())
	_, err = u.Get(id)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSearch(t *testing.T) {
	sort := "updated_at"
	order := "desc"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock_clients.NewMockComputerClient(ctrl)

	var err error
	r := mock_repositories.NewMockComputerRepository(ctrl)
	r.EXPECT().Find(gomock.Any(), gomock.Any(), &database.Options{Sort: sort, Order: order}).Return(err)
	u := NewComputerUsecase(c, r, database.Get())
	_, err = u.Search(&requests.SearchComputersQuery{Sort: &sort, Order: &order})
	if err != nil {
		t.Errorf(err.Error())
	}
}
