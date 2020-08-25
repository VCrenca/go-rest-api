package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"vcrenca/go-rest-api/src/handlers"
	"vcrenca/go-rest-api/src/model/dto"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	hostname     = "localhost"
	port         = 5432
	user         = "postgres"
	password     = "postgres"
	databasename = "go_db"
)

func main() {

	connString := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", port, hostname, user, password, databasename)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize Gin router with default Logger and Recovery system
	r := gin.New()

	r.Use(gin.Logger())

	// Adding the basic recovery middleware with
	r.Use(gin.CustomRecoveryWithWriter(os.Stderr, func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "An error occured"})
		c.Next()
	}))

	// Configure routes
	handlers.ConfigureUserHandler(r, db)

	r.Run(":8080")
}
