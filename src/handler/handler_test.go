package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"pro-backend-trainee-assignment/src/handler"
	"pro-backend-trainee-assignment/src/models"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Service
type MockService struct {
	mock.Mock
}

func (m *MockService) Retrieve(req models.RetrieveRequest) (models.Response, error) {
	args := m.Called(req)
	return args.Get(0).(models.Response), args.Error(1)
}

func (m *MockService) GenerateNumber(req models.GenRequest) (models.Response, error) {
	args := m.Called(req)
	return args.Get(0).(models.Response), args.Error(1)
}
type MockPublisher struct {
	mock.Mock
}

func (mp *MockPublisher) PublishGenerateValue(message []byte) error {
	args := mp.Called(message)
	return args.Error(0)
}

func TestGenerateHandler_Success(t *testing.T) {
	mockPublisher := new(MockPublisher)
	mockPublisher.On("PublishGenerateValue", mock.Anything).Return(nil)

	h := handler.NewHandler(nil, mockPublisher)
	reqBody := `{"type":"string"}`

	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewBufferString(reqBody))
	req.Header.Set("User-Agent", "test-agent")
	w := httptest.NewRecorder()

	h.GenerateHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockPublisher.AssertCalled(t, "PublishGenerateValue", mock.Anything)

}


func TestGenerateHandler_InvalidMethod(t *testing.T) {
	h := handler.NewHandler(nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/generate", nil)
	w := httptest.NewRecorder()

	h.GenerateHandler(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, "method not allowed\n", w.Body.String())
}

func TestGenerateHandler_InvalidType(t *testing.T) {
	h := handler.NewHandler(nil, nil)

	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewBufferString(`{"type":"invalid"}`))
	w := httptest.NewRecorder()

	h.GenerateHandler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "invalid type\n", w.Body.String())
}

func TestRetriveHandler_Success(t *testing.T) {
	mockService := new(MockService)
	mockService.On("Retrieve", mock.Anything).Return(models.Response{}, nil)

	h := handler.NewHandler(mockService, nil)
	req := httptest.NewRequest(http.MethodGet, "/retrieve/123", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	w := httptest.NewRecorder()

	h.RetrieveHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}