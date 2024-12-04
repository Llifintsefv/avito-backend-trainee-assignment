package service

import (
	"fmt"
	"pro-backend-trainee-assignment/src/models"
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

func (m *MockRepository) GetCountRequest(requestID int) (int, error) {
	args := m.Called(requestID)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) UpdateCountRequestAndRetrieveId(requestID int, count int) (string, error) {
	args := m.Called(requestID,count)
	return args.String(0), args.Error(1)
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
// func TestGenerateNumber_SecondRequest(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	serv := NewService(mockRepo)

// 	req := models.GenRequest{
// 		ID:      "test-id",
// 		Type:    "number",
// 		Length:  5,
// 		Values:  nil, 
// 		Url:      "test-url",
// 		UserAgent: "test-user-agent",
// 		RequestId: 1,
// 	}

//     mockRepo.On("GetCountRequest", 1).Return(1, nil)
//     mockRepo.On("UpdateCountRequestAndRetrieveId", "test-request-id",2).Return("test-id", nil)
//     mockRepo.On("Retrieve","test-id").Return("12345","number",nil)

// 	resp, err := serv.GenerateNumber(req)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "test-id", resp.Id)
// 	assert.Equal(t, 12345,resp.Value)
//     mockRepo.AssertExpectations(t)
// }

