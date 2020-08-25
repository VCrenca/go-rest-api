package dal

import (
	"database/sql"
	"vcrenca/go-rest-api/src/model"
)

// IUserRepository -
type IUserRepository interface {
	SaveUser(user model.User) error
	FindByID(id string) (string, error)
	FindAllUsers() ([]string, error)
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
func (dao userRepository) SaveUser(user model.User) error {
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

func (dao userRepository) FindAllUsers() ([]string, error) {
	var userList []string
	sql := "SELECT email FROM users"
	rows, err := dao.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var email string
	for rows.Next() {
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		userList = append(userList, email)
	}

	return userList, nil
}
