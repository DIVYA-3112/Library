package services

import (
	"fmt"

	"github.com/divyathakkar3112/Library/models"
	"github.com/divyathakkar3112/Library/repository"
)

var users []models.User

func CreateUser(user models.User) bool {
	// insert user
	isUserInserted, err := repository.InsertUser(user)
	if !isUserInserted || err != nil {
		fmt.Printf("User Insert Failed")
		return false
	} 
	fmt.Println("User successfully inserted")
	return true
}
