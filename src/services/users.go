package services

import (
	"log"
	"vcrenca/go-rest-api/src/dal"
	"vcrenca/go-rest-api/src/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// IUserService -
type IUserService interface {
	FindByID(id string) (string, error)
	CreateUser(email string, password string) (string, error)
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
		log.Println("Failed getting the user by id !")
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

	user := model.User{
		ID:       uuid,
		Email:    email,
		Password: hash,
	}

	err = svc.repository.SaveUser(user)
	if err != nil {
		log.Println("Failed when creating a user !", err.Error())
		return "", err
	}

	return uuid, nil
}

func encodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println("Failed to encode the password")
		return "", nil
	}
	return string(hash), nil
}
