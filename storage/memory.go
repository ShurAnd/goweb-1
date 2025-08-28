package storage

import (
	"errors"
	"gintest/models"
	"sync"
	"time"
)

type RecipeStorage interface {
	GetAll() []models.Recipe
	GetById(id int) (*models.Recipe, error)
	Create(recipe *models.Recipe) error
}

type InMemoryRecipeStorage struct {
	recipes []models.Recipe
	mu      sync.RWMutex
}

var ErrRecipeNotFound = errors.New("Recipe Not Found")

func InitInMemoryRecipeStorage() *InMemoryRecipeStorage {
	storage := &InMemoryRecipeStorage{
		recipes: []models.Recipe{models.Recipe{
			ID:           1,
			Name:         "Soup",
			Tags:         []string{"soup", "pervoe"},
			Ingredients:  []string{"water", "chicken"},
			Instructions: []string{"get water", "put chicken", "get fire"},
			PublishedAt:  time.Now(),
		}, models.Recipe{
			ID:           2,
			Name:         "Salad",
			Tags:         []string{"salad", "zakuska"},
			Ingredients:  []string{"cucumber", "tomato"},
			Instructions: []string{"cut cucumber", "cut tomato", "mess"},
			PublishedAt:  time.Now(),
		}},
	}

	return storage
}

func (s *InMemoryRecipeStorage) GetAll() []models.Recipe {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.recipes
}
func (s *InMemoryRecipeStorage) GetById(id int) (*models.Recipe, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, r := range s.recipes {
		if r.ID == id {
			return &r, nil
		}
	}

	return nil, ErrRecipeNotFound
}
func (s *InMemoryRecipeStorage) Create(recipe *models.Recipe) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := getNextID(s.recipes)
	recipe.ID = id
	recipe.PublishedAt = time.Now()
	s.recipes = append(s.recipes, *recipe)

	return nil
}

func getNextID(recipes []models.Recipe) int {
	result := 0
	if len(recipes) == 0 {
		return result
	}
	result = recipes[0].ID
	for _, item := range recipes {
		if item.ID > result {
			result = item.ID
		}
	}
	return result + 1
}
