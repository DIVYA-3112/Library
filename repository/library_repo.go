package repository

import (
	"fmt"
	"log"

	"github.com/divyathakkar3112/Library/db"
	"github.com/divyathakkar3112/Library/models"
)

const (
	InsertUserQuery  = `insert into Users (name, email) values (?, ?)`
	GetAllUsersQuery = `select * from Users`
)

func InsertUser(user models.User) (bool, error) {

	// prepare query
	stmt, err := db.DB.Prepare(InsertUserQuery)
	if err != nil {
		log.Printf("Prepare Error - %v", err.Error())
		return false, err
	}
	defer stmt.Close()

	// execute query
	res, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		log.Printf("Exec Error - %v", err.Error())
		return false, err
	}

	lastInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()

	result := fmt.Sprintf("Last inserted - %v \nRows Affected - %v", lastInserted, rowsAffected)
	fmt.Println(result)
	return true, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	stmt, err := db.DB.Prepare(GetAllUsersQuery)
	if err != nil {
		log.Printf("Prepare Error - %v", err.Error())
		return users, err
	}
	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		log.Printf("Execution Error - %v", err.Error())
		return users, err
	}
	defer res.Close()

	for res.Next() {
		var user models.User
		err := res.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Printf("Scanning Error - %v", err.Error())
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
