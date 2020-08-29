package services

import (
	"github.com/vcrenca/go-rest-api/src/dal"
	"github.com/vcrenca/go-rest-api/src/models"
	"github.com/vcrenca/go-rest-api/src/models/dto"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// IUserService -
type IUserService interface {
	FindByID(id string) (string, error)
	CreateUser(email string, password string) (string, error)
	FindAllUsers() ([]dto.GetUserByIDResponse, error)
}

type userService struct {
	repository dal.IUserRepository
}

// NewUserService -
func NewUserService(repo dal.IUserRepository) IUserService {
	return &userService{
		repository: repo,
	}
}

// FindByID -
func (svc userService) FindByID(id string) (string, error) {
	email, err := svc.repository.FindByID(id)
	if err != nil {
		return "", err
	}

	return email, nil
}

// FindByID -
func (svc userService) CreateUser(email string, password string) (string, error) {

	hash, err := encodePassword(password)
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()

	user := models.User{
		ID:       uuid,
		Email:    email,
		Password: hash,
	}

	if err = svc.repository.SaveUser(user); err != nil {
		return "", err
	}

	return uuid, nil
}

func (svc userService) FindAllUsers() ([]dto.GetUserByIDResponse, error) {
	userList, err := svc.repository.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func encodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
