package main

import (
	"database/sql"
	"fmt"

	"github.com/vcrenca/go-rest-api/auth"
	"github.com/vcrenca/go-rest-api/dal"
	"github.com/vcrenca/go-rest-api/handlers"
	"github.com/vcrenca/go-rest-api/server"
	"github.com/vcrenca/go-rest-api/services"
	"github.com/vcrenca/go-rest-api/websocket"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// DB Constants
const (
	hostname     = "postgres-go"
	port         = 5432
	user         = "postgres"
	password     = "postgres"
	databasename = "go_db"
)

func main() {

	// DB Initialization
	connString := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", port, hostname, user, password, databasename)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Server Initialization
	ginServer := server.NewGinServer()

	// Adding logger middleware
	ginServer.Router().Use(gin.Logger())

	// Adding the basic recovery middleware with
	ginServer.Router().Use(gin.RecoveryWithWriter(os.Stderr))

	ginServer.SetPublicGroup("/api")
	ginServer.SetPrivateGroup("/api")

	// Adding the authentication JWT middlware
	ginServer.PrivateGroup().Use(auth.CheckTokenMiddleware)

	// Initiate Repositories
	userRepository := dal.NewUserAccessObject(db)

	// Initiate Services
	userService := services.NewUserService(userRepository)
	authenticationService := auth.NewAuthenticationService(userRepository)

	// Configure routes
	handlers.ConfigureAuthenticationHandler(ginServer, authenticationService)
	handlers.ConfigureUserHandler(ginServer, userService)

	// Websocket
	hub := websocket.NewHub()
	go hub.Run()

	ginServer.Router().GET("/ws", hub.Handler)

	// Launch Server
	ginServer.Router().Run(":8080")
}
