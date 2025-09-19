package repository

import (
	"fmt"
	"log"

	"github.com/divyathakkar3112/Library/db"
	"github.com/divyathakkar3112/Library/models"
)

const (
	InsertUserQuery = `insert into Users (name, email) values (?, ?)`
)

func InsertUser(user models.User) (bool, error) {
	// create db conn
	db.Init()

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
