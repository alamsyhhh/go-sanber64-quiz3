package books

import (
	"errors"
	"time"

	"go-sanber64-quiz3/utils"
)

type BookService interface {
	CreateBook(title, description string, image []byte, releaseYear, price, totalPage, categoryID int, createdBy string) error
	GetBookByID(id int) (*Book, error)
	GetAllBooks() ([]Book, error)
	UpdateBook(id int, title, description string, image []byte, releaseYear, price, totalPage, categoryID int, modifiedBy string) error
	DeleteBook(id int) error
}

type bookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) CreateBook(title, description string, image []byte, releaseYear, price, totalPage, categoryID int, createdBy string) error {
	if releaseYear < 1980 || releaseYear > 2024 {
		return errors.New("tahun rilis harus antara 1980 - 2024")
	}

	imageURL, err := utils.UploadImageToCloudinary(image)
	if err != nil {
		return err
	}

	thickness := "tipis"
	if totalPage > 100 {
		thickness = "tebal"
	}

	book := &Book{
		Title:       title,
		Description: description,
		ImageURL:    imageURL,
		ReleaseYear: releaseYear,
		Price:       price,
		TotalPage:   totalPage,
		Thickness:   thickness,
		CategoryID:  categoryID,
		CreatedAt:   time.Now(),
		CreatedBy:   createdBy,
		ModifiedAt:  time.Now(),
		ModifiedBy:  createdBy,
	}
	return s.repo.CreateBook(book)
}

func (s *bookService) GetBookByID(id int) (*Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *bookService) GetAllBooks() ([]Book, error) {
	return s.repo.GetAllBooks()
}

func (s *bookService) UpdateBook(id int, title, description string, image []byte, releaseYear, price, totalPage, categoryID int, modifiedBy string) error {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return errors.New("buku tidak ditemukan")
	}

	if len(image) > 0 {
		imageURL, err := utils.UploadImageToCloudinary(image)
		if err != nil {
			return err
		}
		book.ImageURL = imageURL
	}

	book.Title = title
	book.Description = description
	book.ReleaseYear = releaseYear
	book.Price = price
	book.TotalPage = totalPage
	book.Thickness = "tipis"
	if totalPage > 100 {
		book.Thickness = "tebal"
	}
	book.CategoryID = categoryID
	book.ModifiedAt = time.Now()
	book.ModifiedBy = modifiedBy

	return s.repo.UpdateBook(book)
}

func (s *bookService) DeleteBook(id int) error {
	_, err := s.repo.GetBookByID(id)
	if err != nil {
		return errors.New("buku tidak ditemukan")
	}

	return s.repo.DeleteBook(id)
}
