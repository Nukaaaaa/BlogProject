package mock

import (
	"myproject/models"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

// --- Методы для постов (у тебя уже есть) ---
func (m *MockDB) FindPosts(userID uint) ([]models.Post, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Post), args.Error(1)
}

func (m *MockDB) CreatePost(post models.Post) (*models.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockDB) FindPostByID(id uint) (*models.Post, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockDB) UpdatePost(id uint, post models.Post) (*models.Post, error) {
	args := m.Called(id, post)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockDB) DeletePost(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDB) FindPostsByCategory(categoryID uint) ([]models.Post, error) {
	args := m.Called(categoryID)
	return args.Get(0).([]models.Post), args.Error(1)
}

// --- Добавляем методы для категорий ---
func (m *MockDB) FindCategories() ([]models.Category, error) {
	args := m.Called()
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockDB) CreateCategory(cat models.Category) (*models.Category, error) {
	args := m.Called(cat)
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockDB) FindCategoryByID(id uint) (*models.Category, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockDB) UpdateCategory(id uint, cat models.Category) (*models.Category, error) {
	args := m.Called(id, cat)
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockDB) DeleteCategory(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
