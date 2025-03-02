package books

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
)

type BookRepository interface {
	CreateBook(book *Book) error
	GetBookByID(id int) (*Book, error)
	GetAllBooks() ([]Book, error)
	UpdateBook(book *Book) error
	DeleteBook(id int) error
}

type bookRepository struct {
	db *goqu.Database
}

func NewBookRepository(db *goqu.Database) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) CreateBook(book *Book) error {
	query := r.db.Insert("books").
		Cols("title", "description", "image_url", "release_year", "price", "total_page", "thickness", "category_id", "created_at", "created_by", "modified_at", "modified_by").
		Vals(goqu.Vals{book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedAt, book.CreatedBy, book.ModifiedAt, book.ModifiedBy}).
		Returning("id")

	var newID int
	_, err := query.Executor().ScanVal(&newID)
	if err != nil {
		return err
	}
	book.ID = newID
	return nil
}

func (r *bookRepository) GetBookByID(id int) (*Book, error) {
	var book Book
	found, err := r.db.From("books").Where(goqu.Ex{"id": id}).ScanStruct(&book)
	if err != nil || !found {
		return nil, errors.New("buku tidak ditemukan")
	}
	return &book, nil
}

func (r *bookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	err := r.db.From("books").ScanStructs(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) UpdateBook(book *Book) error {
	_, err := r.db.Update("books").
		Set(goqu.Record{
			"title":       book.Title,
			"description": book.Description,
			"image_url":   book.ImageURL,
			"release_year": book.ReleaseYear,
			"price":       book.Price,
			"total_page":  book.TotalPage,
			"thickness":   book.Thickness,
			"category_id": book.CategoryID,
			"modified_at": book.ModifiedAt,
			"modified_by": book.ModifiedBy,
		}).
		Where(goqu.Ex{"id": book.ID}).
		Executor().Exec()
	return err
}

func (r *bookRepository) DeleteBook(id int) error {
	_, err := r.db.Delete("books").
		Where(goqu.Ex{"id": id}).
		Executor().Exec()
	return err
}
