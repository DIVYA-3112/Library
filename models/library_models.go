package models

// type Books struct {
// 	bookName string
// 	bookId   int64
// 	quantity int64
// }

type User struct {
	Name     string `json:"name"`
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}
