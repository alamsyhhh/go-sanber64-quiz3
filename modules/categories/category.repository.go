package categories

import (
	"database/sql"
	"errors"
	"fmt"

	"go-sanber64-quiz3/modules/books"

	"github.com/doug-martin/goqu/v9"
)

type CategoryRepository interface {
	CreateCategory(category *Category) error
	GetCategoryByID(id int) (*Category, error)
	GetAllCategories() ([]Category, error)
	UpdateCategory(category *Category) error
	DeleteCategory(id int) error
	GetBooksByCategory(categoryID int) ([]books.Book, error)
}

type categoryRepository struct {
	db *goqu.Database
}

func NewCategoryRepository(db *goqu.Database) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *Category) error {
	query := r.db.Insert("categories").
		Cols("name", "created_at", "created_by", "modified_at", "modified_by").
		Vals(goqu.Vals{category.Name, category.CreatedAt, category.CreatedBy, category.ModifiedAt, category.ModifiedBy}).
		Returning("id")
			
	var newID int
	_, err := query.Executor().ScanVal(&newID)
	if err != nil {
		return err
	}
	category.ID = newID
	return nil
}

func (r *categoryRepository) GetCategoryByID(id int) (*Category, error) {
	var category Category
	found, err := r.db.From("categories").Where(goqu.Ex{"id": id}).ScanStruct(&category)
	if err != nil || !found {
		return nil, errors.New("kategori tidak ditemukan")
	}
	return &category, nil
}

func (r *categoryRepository) GetAllCategories() ([]Category, error) {
	var categories []Category
	err := r.db.From("categories").ScanStructs(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) UpdateCategory(category *Category) error {
	_, err := r.db.Update("categories").
		Set(goqu.Record{
			"name":        category.Name,
			"modified_at": category.ModifiedAt,
			"modified_by": category.ModifiedBy,
		}).
		Where(goqu.Ex{"id": category.ID}).
		Executor().Exec()
	return err
}

func (r *categoryRepository) DeleteCategory(id int) error {
	_, err := r.db.Delete("categories").
		Where(goqu.Ex{"id": id}).
		Executor().Exec()
	return err
}

func (r *categoryRepository) GetBooksByCategory(categoryID int) ([]books.Book, error) {
	var books []books.Book
	err := r.db.From("books").Where(goqu.Ex{"category_id": categoryID}).ScanStructs(&books)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("gagal mengambil buku: %w", err)
	}

	return books, nil
}

