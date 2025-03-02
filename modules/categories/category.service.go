package categories

import (
	"errors"
	"go-sanber64-quiz3/modules/books"
	"time"
)

type CategoryService interface {
	CreateCategory(name, createdBy string) (*Category, error)
	GetCategoryByID(id int) (*Category, error)
	GetAllCategories() ([]Category, error)
	UpdateCategory(id int, name, modifiedBy string) (*Category, error)
	DeleteCategory(id int) (*Category, error)
	GetBooksByCategory(categoryID int) ([]books.Book, error)
}

type categoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategory(name, createdBy string) (*Category, error) {
	category := &Category{
		Name:       name,
		CreatedAt:  time.Now(),
		CreatedBy:  createdBy,
		ModifiedAt: time.Now(),
		ModifiedBy: createdBy,
	}
	err := s.repo.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}


func (s *categoryService) GetCategoryByID(id int) (*Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *categoryService) GetAllCategories() ([]Category, error) {
	return s.repo.GetAllCategories()
}

func (s *categoryService) UpdateCategory(id int, name, modifiedBy string) (*Category, error) {
	category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = name
	category.ModifiedAt = time.Now()
	category.ModifiedBy = modifiedBy

	updatedCategory, err := s.repo.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (s *categoryService) DeleteCategory(id int) (*Category, error) {
	_, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}

	deletedCategory, err := s.repo.DeleteCategory(id)
	if err != nil {
		return nil, err
	}

	return deletedCategory, nil
}


func (s *categoryService) GetBooksByCategory(categoryID int) ([]books.Book, error) {
	return s.repo.GetBooksByCategory(categoryID)
}
