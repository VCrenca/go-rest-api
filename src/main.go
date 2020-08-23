package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"vcrenca/go-rest-api/src/dal"
	"vcrenca/go-rest-api/src/handlers"
	"vcrenca/go-rest-api/src/services"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	userRepository := dal.NewUserAccessObject(db)
	userService := services.NewUserService(userRepository)

	handlers.ConfigureUserHandler(r, userService)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Starting the server on port 8080...")
	log.Fatal(server.ListenAndServe())
}
