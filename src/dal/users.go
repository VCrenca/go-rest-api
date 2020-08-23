package dal

import (
	"database/sql"
	"log"
	"vcrenca/go-rest-api/src/model"
)

// IUserRepository -
type IUserRepository interface {
	SaveUser(user model.User) error
	FindByID(id string) (string, error)
}

// userRepository -
type userRepository struct {
	db *sql.DB
}

// NewUserAccessObject -
func NewUserAccessObject(db *sql.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

// SaveUser -
func (dao *userRepository) SaveUser(user model.User) error {
	sql := "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)"
	_, err := dao.db.Exec(sql, user.ID, user.Email, user.Password)
	if err != nil {
		log.Println("Error while saving a user !")
	}

	return nil
}

// GetUserByID -
func (dao *userRepository) FindByID(id string) (string, error) {
	var email string
	sql := "SELECT email FROM users WHERE id = $1"
	row := dao.db.QueryRow(sql, id)
	err := row.Scan(&email)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return email, nil
}
