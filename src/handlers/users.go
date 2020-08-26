package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"vcrenca/go-rest-api/src/dal"
	"vcrenca/go-rest-api/src/models/dto"
	"vcrenca/go-rest-api/src/server"
	"vcrenca/go-rest-api/src/services"

	"github.com/gin-gonic/gin"
)

// UserHandler -
type UserHandler struct {
	svc services.IUserService
}

// ConfigureUserHandler -
func ConfigureUserHandler(ginServer *server.GinServer, db *sql.DB) {

	userRepository := dal.NewUserAccessObject(db)
	userService := services.NewUserService(userRepository)

	handler := UserHandler{
		svc: userService,
	}

	// Private routes
	ginServer.PrivateGroup().GET("/users", handler.GetAllUsers)
	ginServer.PrivateGroup().GET("/users/:id", handler.GetUserByID)

	// Public routes
	ginServer.PublicGroup().POST("/users", handler.PostUser)
}

// GetUserByID -
func (h UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	email, err := h.svc.FindByID(id)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, dto.GetUserByIDResponse{Email: email})
}

// PostUser -
func (h UserHandler) PostUser(c *gin.Context) {
	var userRequest dto.CreateUserRequest

	err := c.BindJSON(&userRequest)
	if err != nil {
		panic(err.Error())
	}

	if userRequest.Email == "" || userRequest.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: "You need to provide an email and a password !"})
	}

	id, err := h.svc.CreateUser(userRequest.Email, userRequest.Password)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Created user :", id)
	c.JSON(http.StatusOK, dto.CreateUserResponse{ID: id})
}

// GetAllUsers -
func (h UserHandler) GetAllUsers(c *gin.Context) {

	var response []dto.GetUserByIDResponse

	userList, err := h.svc.FindAllUsers()
	if err != nil {
		panic(err.Error())
	}

	for _, email := range userList {
		response = append(response, dto.GetUserByIDResponse{Email: email})
	}

	c.JSON(http.StatusOK, response)
}
