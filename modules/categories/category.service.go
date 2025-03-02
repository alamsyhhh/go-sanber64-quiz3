package categories

import (
	"errors"
	"go-sanber64-quiz3/modules/books"
	"time"
)

type CategoryService interface {
	CreateCategory(name, createdBy string) error
	GetCategoryByID(id int) (*Category, error)
	GetAllCategories() ([]Category, error)
	UpdateCategory(id int, name, modifiedBy string) error
	DeleteCategory(id int) error
	GetBooksByCategory(categoryID int) ([]books.Book, error)
}

type categoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategory(name, createdBy string) error {
	
	category := &Category{
		Name:       name,
		CreatedAt:  time.Now(),
		CreatedBy:  createdBy,
		ModifiedAt: time.Now(),
		ModifiedBy: createdBy,
	}
	return s.repo.CreateCategory(category)
}

func (s *categoryService) GetCategoryByID(id int) (*Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *categoryService) GetAllCategories() ([]Category, error) {
	return s.repo.GetAllCategories()
}

func (s *categoryService) UpdateCategory(id int, name, modifiedBy string) error {
	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return err
	}

	category.Name = name
	category.ModifiedAt = time.Now()
	category.ModifiedBy = modifiedBy

	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id int) error {
	_, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return errors.New("kategori tidak ditemukan")
	}
	return s.repo.DeleteCategory(id)
}

func (s *categoryService) GetBooksByCategory(categoryID int) ([]books.Book, error) {
	return s.repo.GetBooksByCategory(categoryID)
}
