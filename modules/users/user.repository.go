package users

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) error
}

type userRepository struct {
	db *goqu.Database
}

func NewUserRepository(db *goqu.Database) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *User) error {
    query := r.db.Insert("users").
        Cols("username", "password", "created_at", "created_by", "modified_at", "modified_by").
        Vals(goqu.Vals{user.Username, user.Password, user.CreatedAt, user.CreatedBy, user.ModifiedAt, user.ModifiedBy}).
        Returning("id")

    var newID int
	_, err := query.Executor().ScanVal(&newID)
	if err != nil {
		return err
	}
	user.ID = newID

    return nil
}

func (r *userRepository) GetUserByUsername(username string) (*User, error) {
	var user User
	found, err := r.db.From("users").Where(goqu.Ex{"username": username}).ScanStruct(&user)
	if err != nil || !found {
		return nil, errors.New("user tidak ditemukan")
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *User) error {
	_, err := r.db.Update("users").
		Set(goqu.Record{
			"username":    user.Username,
			"modified_at": user.ModifiedAt,
			"modified_by": user.ModifiedBy,
		}).
		Where(goqu.Ex{"id": user.ID}).
		Executor().Exec()
	return err
}

func (r *userRepository) GetUserByID(id int) (*User, error) {
	var user User
	found, err := r.db.From("users").Where(goqu.Ex{"id": id}).ScanStruct(&user)
	if err != nil || !found {
		return nil, errors.New("user tidak ditemukan")
	}
	return &user, nil
}
