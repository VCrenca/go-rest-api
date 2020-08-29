package dal

import (
	"database/sql"

	"github.com/vcrenca/go-rest-api/models/dto"

	"github.com/vcrenca/go-rest-api/models"
)

// IUserRepository -
type IUserRepository interface {
	SaveUser(user models.User) error
	FindByID(id string) (string, error)
	FindAllUsers() ([]dto.UserResponse, error)
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
func (dao userRepository) SaveUser(user models.User) error {
	sql := "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)"
	_, err := dao.db.Exec(sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID -
func (dao userRepository) FindByID(id string) (string, error) {
	var email string
	sql := "SELECT email FROM users WHERE id = $1"
	row := dao.db.QueryRow(sql, id)
	err := row.Scan(&email)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (dao userRepository) FindAllUsers() ([]dto.UserResponse, error) {
	var userList []dto.UserResponse
	sql := "SELECT id, email FROM users"
	rows, err := dao.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id string
	var email string
	for rows.Next() {
		if err := rows.Scan(&id, &email); err != nil {
			return nil, err
		}
		userList = append(userList, dto.UserResponse{ID: id, Email: email})
	}

	return userList, nil
}
