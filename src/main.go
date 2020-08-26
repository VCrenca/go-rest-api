package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"vcrenca/go-rest-api/src/handlers"
	"vcrenca/go-rest-api/src/models/dto"
	"vcrenca/go-rest-api/src/server"

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
	ginServer.Router().Use(gin.CustomRecoveryWithWriter(os.Stderr, func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "An error occured"})
		c.Next()
	}))

	ginServer.SetPublicGroup("/api")
	ginServer.SetPrivateGroup("/api")

	// Configure routes
	handlers.ConfigureUserHandler(ginServer, db)

	ginServer.Router().Run(":8080")
}
