package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vcrenca/go-rest-api/dal"
	"github.com/vcrenca/go-rest-api/models"
	"github.com/vcrenca/go-rest-api/models/dto"
	"golang.org/x/crypto/bcrypt"
)

// IAuthenticationService -
type IAuthenticationService interface {
	CheckEmailAndPassword(email string, password string) (string, error)
}

type authenticationService struct {
	userRepository dal.IUserRepository
}

// NewAuthenticationService  -
func NewAuthenticationService(userRepository dal.IUserRepository) IAuthenticationService {
	return &authenticationService{
		userRepository: userRepository,
	}
}

// CheckTokenMiddleware -
func CheckTokenMiddleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization != "" {
		token := strings.Split(authorization, " ")[1]
		if err := checkToken(token); err != nil {
			log.Println("Token is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Unauthorized"})
		}
	} else {
		log.Println("No authorization header found !")
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Unauthorized"})
	}
}

func checkPassword(password string, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}

	return nil
}

func createToken(user models.User) (string, error) {

	now := time.Now()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: now.Add(10 * time.Hour).Unix(),
		IssuedAt:  now.Unix(),
		Subject:   user.ID,
	})

	token, err := newToken.SignedString([]byte("verybigsecret"))
	if err != nil {
		panic(err)
	}

	return token, nil
}

// Check the email and password passed for login and return the token if it is correct, otherwise return an error
func (a authenticationService) CheckEmailAndPassword(email string, password string) (string, error) {
	user, err := a.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = checkPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := createToken(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func checkToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("verybigsecret"), nil
	})

	return err
}
