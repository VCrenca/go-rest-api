package main

import (
	"database/sql"
	"fmt"

	"github.com/vcrenca/go-rest-api/dal"
	"github.com/vcrenca/go-rest-api/handlers"
	"github.com/vcrenca/go-rest-api/server"
	"github.com/vcrenca/go-rest-api/services"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// DB Constants
const (
	hostname     = "localhost"
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

	// Initiate Repositories
	userRepository := dal.NewUserAccessObject(db)

	// Initiate Services
	userService := services.NewUserService(userRepository)

	// Configure routes
	handlers.ConfigureUserHandler(ginServer, userService)

	// Launch Server
	ginServer.Router().Run(":8080")
}
