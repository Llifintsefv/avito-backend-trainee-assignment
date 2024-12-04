package service

import (
	"errors"
	"fmt"
	"pro-backend-trainee-assignment/src/models"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Generate(value models.GenerateValue) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockRepository) Retrieve(id string) (string, string, error) {
	args := m.Called(id)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockRepository) GetCountRequest(requestId int) (int, error) {
	args := m.Called(requestId)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) UpdateCountRequestAndRetrieveId(requestId int, countRequest int) (string, error) {
	args := m.Called(requestId, countRequest)
	return args.String(0), args.Error(1)
}

func TestNewService(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)
	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
}

func TestService_Retrieve(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	req := models.RetrieveRequest{
		ID:        "test_id",
		UserAgent: "test_ua",
		Url:       "test_url",
		RequestId: 123,
	}

	mockRepo.On("Retrieve", "test_id").Return("test_value", "string", nil)
	mockRepo.On("Generate", mock.AnythingOfType("models.GenerateValue")).Return(nil)

	response, err := service.Retrieve(req)

	assert.NoError(t, err)
	assert.Equal(t, "test_id", response.Id)
	assert.Equal(t, "test_value", response.Value)

	expectedGenValue := models.GenerateValue{
		ID:          "test_id",
		UserAgent:   "test_ua",
		Url:         "test_url",
		RequestId:   123,
	}

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Retrieve", "test_id")
	mockRepo.AssertCalled(t, "Generate", mock.MatchedBy(func(gv models.GenerateValue) bool {
		return reflect.DeepEqual(gv, expectedGenValue)
	}))
}

func TestService_Retrieve_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	req := models.RetrieveRequest{
		ID:        "test_id",
		UserAgent: "test_ua",
		Url:       "test_url",
		RequestId:   123,
	}

	mockRepo.On("Retrieve", "test_id").Return("", "", errors.New("retrieve error"))

	_, err := service.Retrieve(req)

	assert.Error(t, err)
	assert.Equal(t, "retrieve error", err.Error())

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Retrieve", "test_id")
    mockRepo.AssertNotCalled(t, "Generate", mock.Anything)
}

func TestService_GetValueById_Number(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockRepo.On("Retrieve", "test_id").Return("123", "number", nil)

	response, err := service.GetValueById("test_id")

	assert.NoError(t, err)
	assert.Equal(t, "test_id", response.Id)
	assert.Equal(t, 123, response.Value)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Retrieve", "test_id")
}

func TestService_GetValueById_String(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockRepo.On("Retrieve", "test_id").Return("test_value", "string", nil)

	response, err := service.GetValueById("test_id")

	assert.NoError(t, err)
	assert.Equal(t, "test_id", response.Id)
	assert.Equal(t, "test_value", response.Value)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Retrieve", "test_id")
}

func TestService_GetValueById_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockRepo.On("Retrieve", "test_id").Return("", "", errors.New("retrieve error"))

	_, err := service.GetValueById("test_id")

	assert.Error(t, err)
	assert.Equal(t, "retrieve error", err.Error())

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Retrieve", "test_id")
}

func TestService_GetValueById_NumberConversionError(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockRepo.On("Retrieve", "test_id").Return("abc", "number", nil)

	_, err := service.GetValueById("test_id")

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Retrieve", "test_id")
}


func TestGenerateNumber_FirstRequest(t *testing.T) {
	mockRepo := new(MockRepository)
	serv := NewService(mockRepo)

	req := models.GenRequest{
		ID:      "test-id",
		Type:    "number",
		Length:  5,
		Values:  nil, 
		Url:      "test-url",
		UserAgent: "test-user-agent",
		RequestId: 1,
	}

    mockRepo.On("GetCountRequest", 1).Return(0, nil)
    mockRepo.On("Generate", mock.AnythingOfType("models.GenerateValue")).Return(nil)

	resp, err := serv.GenerateNumber(req)

	assert.NoError(t, err)
	assert.Equal(t, "test-id", resp.Id)
	assert.Len(t, fmt.Sprint(resp.Value), 5) 
    mockRepo.AssertExpectations(t)
}
