package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vcrenca/go-rest-api/auth"
	"github.com/vcrenca/go-rest-api/models/dto"
	"github.com/vcrenca/go-rest-api/server"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticationHandler -
type AuthenticationHandler struct {
	authService auth.IAuthenticationService
}

// ConfigureAuthenticationHandler -
func ConfigureAuthenticationHandler(ginServer *server.GinServer, svc auth.IAuthenticationService) {

	handler := AuthenticationHandler{
		authService: svc,
	}

	// Public routes
	ginServer.PublicGroup().POST("/login", handler.Login)
}

// Login -
func (h AuthenticationHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	c.BindJSON(&loginRequest)
	token, err := h.authService.CheckEmailAndPassword(loginRequest.Email, loginRequest.Password)
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword, sql.ErrNoRows:
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Email/Password doesn't exists, please try again !"})
			return
		default:
			panic(err)
		}
	}

	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Status(http.StatusOK)
}
