package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrs24/pkg/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var NewUser models.User

type User struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	Role           string   `json:"role"`
	Parent         string   `json:"parent"`
	JGroup         []string `json:"jgroup"`
	Permission     string   `json:"permission"`
	Requests       string   `json:"requests"`
	EditPermission string   `json:"editPermission"`
}

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	newBooks := models.GetAllBooks()
// 	res, _ := json.Marshal(newBooks)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetBookById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	bookDetails, _ := models.GetBookById(ID)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

//	func CreateBook(w http.ResponseWriter, r *http.Request) {
//		CreateBook := &models.Book{}
//		utils.ParseBody(r, CreateBook)
//		b := CreateBook.CreateBook()
//		res, _ := json.Marshal(b)
//		w.WriteHeader(http.StatusOK)
//		w.Write(res)
//	}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user User

	// Decode the JSON data from the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle the error, if any
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	db, err := sql.Open("mysql", "root:root@8085@tcp(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// The values to be inserted
	name := user.Name
	email := user.Email
	role := user.Role
	parent := user.Parent
	jgroup := `["group1", "group2", "group3"]`
	permission := user.Permission
	requests := user.Requests
	editPermission := user.EditPermission

	// The insert query with prepared statement
	query := `INSERT INTO users (name, email, role, parent, jgroup, permission, requests, editPermission) VALUES (?, ?, ?,  ?, ?, ?, ?, ?)`

	// Execute the query with prepared statement
	result, err := db.Exec(query, name, email, role, parent, jgroup, permission, requests, editPermission)
	if err != nil {
		panic(err)
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Rows affected:", rowsAffected)

}

// responseData := struct {
// 	Name string `json:"name"`
// }{
// 	Name: user.Name,
// }

// // Set the appropriate content type for the response
// w.Header().Set("Content-Type", "application/json")

// // Encode the response data and send it back
// err = json.NewEncoder(w).Encode(responseData)
// if err != nil {
// 	// Handle the error, if any
// 	http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
// 	return
// }

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	book := models.DeleteBook(ID)
// 	res, _ := json.Marshal(book)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	var updateBook = &models.Book{}
// 	utils.ParseBody(r, updateBook)
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	bookDetails, db := models.GetBookById(ID)
// 	if updateBook.Name != "" {
// 		bookDetails.Name = updateBook.Name
// 	}
// 	if updateBook.Author != "" {
// 		bookDetails.Author = updateBook.Author
// 	}
// 	if updateBook.Publication != "" {
// 		bookDetails.Publication = updateBook.Publication
// 	}
// 	db.Save(&bookDetails)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
