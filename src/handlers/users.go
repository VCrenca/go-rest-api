package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vcrenca/go-rest-api/src/model/dto"
	"vcrenca/go-rest-api/src/services"

	"github.com/gorilla/mux"
)

// UserHandler -
type UserHandler struct {
	svc services.IUserService
}

// ConfigureUserHandler -
func ConfigureUserHandler(r *mux.Router, svc services.IUserService) {
	handler := UserHandler{
		svc: svc,
	}
	r.Methods("GET").Path("/users/{id}").HandlerFunc(handler.GetUserByID)
	r.Methods("POST").Path("/users").HandlerFunc(handler.PostUser)
}

// GetUserByID -
func (h UserHandler) GetUserByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if vars["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email, err := h.svc.FindByID(vars["id"])
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Message: "An error occured"})
		return
	}

	log.Println("Retrieving from database :", email)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.GetUserByIDResponse{Email: email})
}

// PostUser -
func (h UserHandler) PostUser(w http.ResponseWriter, req *http.Request) {

	var userRequest dto.CreateUserRequest
	err := json.NewDecoder(req.Body).Decode(&userRequest)
	if err != nil {
		log.Println("Error while decoding request body", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userRequest.Email == "" || userRequest.Password == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorResponse{Message: "You need to provide an email and a password !"})
		return
	}

	id, err := h.svc.CreateUser(userRequest.Email, userRequest.Password)
	if err != nil {
		log.Println("Error while creating the user", err.Error())
		return
	}

	log.Println("Created user :", id)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.CreateUserResponse{ID: id})
}
