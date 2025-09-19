package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/divyathakkar3112/Library/models"
	"github.com/divyathakkar3112/Library/services"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Library")
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var userReq models.User
	var resp models.Response
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}
	isUserCreated := services.CreateUser(userReq)
	if isUserCreated {
		resp = models.Response{
			Status:  http.StatusOK,
			Message: "New user is created",
		}
	} else {
		resp = models.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create a new user",
		}
	}

	w.Header().Set("contents-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(&resp)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var resp models.Response
	userData, err := services.GetAllUsersData()
	if err != nil {
		log.Printf("Error getting all users")
		resp.Status = http.StatusInternalServerError
		resp.Message = "Error getting Data"
		resp.Data = userData
	} else {
		log.Printf("Data fetched successfully")
		resp.Status = http.StatusOK
		resp.Message = "Data fetched successfullly"
		resp.Data = userData
	}

	w.Header().Set("contents-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(&resp)
}
