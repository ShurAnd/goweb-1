package handlers

import (
	"encoding/json"
	"gintest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockStorage для тестов
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetAll() []models.Recipe {
	args := m.Called()
	return args.Get(0).([]models.Recipe)
}

func (m *MockStorage) GetById(id int) (*models.Recipe, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Recipe), args.Error(1)
}

func (m *MockStorage) Create(recipe *models.Recipe) error {
	args := m.Called(recipe)
	return args.Error(0)
}

func (m *MockStorage) GetNextID() int {
	args := m.Called()
	return args.Int(0)
}

func TestGetRecipes(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	mockStorage := new(MockStorage)
	expectedRecipes := []models.Recipe{{ID: 1, Name: "Test"}}
	mockStorage.On("GetAll").Return(expectedRecipes)

	handler := InitRecipeHandler(mockStorage)

	// Test
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler.GetRecipes(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]models.Recipe
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Len(t, response["recipes"], 1)
	assert.Equal(t, "Test", response["recipes"][0].Name)
	mockStorage.AssertExpectations(t)
}
