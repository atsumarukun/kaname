package handlers

import (
	"backend/internal/app/api/computers/domains/entities"
	"backend/internal/app/api/computers/interfaces/requests"
	"backend/internal/app/api/computers/interfaces/responses"
	mock_usecases "backend/internal/app/api/computers/mocks/usecases"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Data *responses.Computer `json:"data"`
}

type SearchResponse struct {
	Data *[]responses.Computer `json:"data"`
}

func TestCreateComputer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	input := requests.CreateComputerInput{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}

	body, err := json.Marshal(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	req, err := http.NewRequest("POST", "/api/v1/computers/", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = req

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	computer := &entities.Computer{
		HostName:   input.HostName,
		IPAddress:  input.IPAddress,
		MACAddress: input.MACAddress,
	}
	u := mock_usecases.NewMockComputerUsecase(ctrl)
	u.EXPECT().Create(input).Return(computer, err)

	h := NewComputerHandler(u)
	h.CreateComputer(context)

	var res Response
	err = json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, res.Data.HostName, input.HostName)
	assert.Equal(t, res.Data.IPAddress, input.IPAddress)
	assert.Equal(t, res.Data.MACAddress, input.MACAddress)
}

func TestUpdateComputer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	id := uint(1)
	input := requests.UpdateComputerInput{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}

	body, err := json.Marshal(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/computers/%d", id), bytes.NewBuffer(body))
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = req
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.FormatUint(uint64(id), 10),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	computer := &entities.Computer{
		HostName:   input.HostName,
		IPAddress:  input.IPAddress,
		MACAddress: input.MACAddress,
	}
	u := mock_usecases.NewMockComputerUsecase(ctrl)
	u.EXPECT().Update(id, input).Return(computer, err)

	h := NewComputerHandler(u)
	h.UpdateComputer(context)

	var res Response
	err = json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, res.Data.HostName, input.HostName)
	assert.Equal(t, res.Data.IPAddress, input.IPAddress)
	assert.Equal(t, res.Data.MACAddress, input.MACAddress)
}

func TestDeleteComputer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	id := uint(1)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/computers/%d", id), nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = req
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.FormatUint(uint64(id), 10),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	computer := &entities.Computer{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}
	u := mock_usecases.NewMockComputerUsecase(ctrl)
	u.EXPECT().Delete(id).Return(computer, err)

	h := NewComputerHandler(u)
	h.DeleteComputer(context)

	var res Response
	err = json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestWakeComputer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	id := uint(1)

	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/computers/%d/wake", id), nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = req
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.FormatUint(uint64(id), 10),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	computer := &entities.Computer{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}
	u := mock_usecases.NewMockComputerUsecase(ctrl)
	u.EXPECT().Wake(id).Return(computer, err)

	h := NewComputerHandler(u)
	h.WakeComputer(context)

	var res Response
	err = json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetComputer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	id := uint(1)

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/computers/%d", id), nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = req
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.FormatUint(uint64(id), 10),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	computer := &entities.Computer{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}
	u := mock_usecases.NewMockComputerUsecase(ctrl)
	u.EXPECT().Get(id).Return(computer, err)

	h := NewComputerHandler(u)
	h.GetComputer(context)

	var res Response
	err = json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestSearchComputer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	query := &requests.SearchComputersQuery{}

	req, err := http.NewRequest("GET", "/api/v1/computers/", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = req

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	computers := &[]entities.Computer{
		{
			HostName:   "Testing",
			IPAddress:  "0.0.0.0",
			MACAddress: "00:00:00:00:00:00",
		},
	}
	u := mock_usecases.NewMockComputerUsecase(ctrl)
	u.EXPECT().Search(query).Return(computers, err)

	h := NewComputerHandler(u)
	h.SearchComputers(context)

	var res SearchResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, w.Code, http.StatusOK)
}
