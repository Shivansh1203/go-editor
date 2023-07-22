package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/go-sql-driver/mysql"

	"github.com/rs/cors"
)

var db *gorm.DB

type User struct {
	Username       string `json:"username"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	Parent         string `json:"parent"`
	JGroup         []byte `json:"author"`
	Permission     string `json:"permission"`
	Requests       string `json:"requests"`
	EditPermission bool   `json:"editpermission"`
}

type Taskcomp struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Parent      string `json:"parent"`
	JInput      string `json:"jInput"`
	Language    string `json:"language"`
	JGroupAcess []byte `json:"jGroupAcess"`
	Flag        bool   `json:"editpermission"`
}

type Userlog struct {
	gorm.Model
	Username string `"json:"username"`
	Password string `"json:"password"`
}
type Usertask struct {
	gorm.Model
	Name  string `"json:"name"`
	Input string `"json:"input"`
}
type Task struct {
	gorm.Model
	Name string `"json:"name"`
}

func Chk(w http.ResponseWriter, r *http.Request) {
	log.Println("GET CHK ")
	fmt.Println("Rows")
}

func getTaks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mysql", "root:Ras$1203@tcp(localhost:3306)/editor")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := "SELECT * FROM TaskLog"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	// var columns []string
	// columns = []string{"hello", "world", "goodbye"}

	for rows.Next() {
		var (
			Name        string
			Type        string
			Parent      string
			jInput      string
			Language    string
			jGroupAcess string
			Flag        string
		)
		rows.Scan()
		if err := rows.Scan(&Name, &Type, &Parent, &jInput, &Language, &jGroupAcess, &Flag); err != nil {
			log.Fatal(err)
		}
		log.Printf(Name, Type, Parent, jInput, Language, jGroupAcess, Flag)
	}

	// fmt.Println("Rows affected:", rowsAffected)

}

func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mysql", "root:Ras$1203@tcp(localhost:3306)/editor")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := "SELECT * FROM TaskLog"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	// var columns []string
	// columns = []string{"hello", "world", "goodbye"}

	for rows.Next() {
		var (
			Name        string
			Type        string
			Parent      string
			jInput      string
			Language    string
			jGroupAcess string
			Flag        string
		)
		rows.Scan()
		if err := rows.Scan(&Name, &Type, &Parent, &jInput, &Language, &jGroupAcess, &Flag); err != nil {
			log.Fatal(err)
		}
		log.Printf(Name, Type, Parent, jInput, Language, jGroupAcess, Flag)
	}

	// fmt.Println("Rows affected:", rowsAffected)

}
func handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user User

	// Decode the JSON data from the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Handle the error, if any
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	db, err := sql.Open("mysql", "root:Ras$1203@tcp(localhost:3306)/editor")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// The values to be inserted
	username := user.Username
	name := user.Name
	password := user.Password
	role := user.Role
	parent := user.Parent
	jgroup := `["group1", "group2", "group3"]`
	permission := user.Permission
	requests := user.Requests
	editPermission := user.EditPermission
	log.Println(username, password, name)

	// The insert query with prepared statement
	query := `INSERT INTO users (username,name,password, role, parent, jgroup, permission, requests, editPermission) VALUES (?, ?,?, ?,  ?, ?, ?, ?, ?)`

	// Execute the query with prepared statement
	result, err := db.Exec(query, username, name, password, role, parent, jgroup, permission, requests, editPermission)
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

func handleCommit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var usertask Usertask

	// Decode the JSON data from the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&usertask)
	if err != nil {
		// Handle the error, if any
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	db, err := sql.Open("mysql", "root:Ras$1203@tcp(localhost:3306)/editor")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// The values to be inserted
	name := usertask.Name
	input := usertask.Input

	fmt.Println("name:", name)
	fmt.Println("input", input)
	// Prepare the UPDATE statement.
	stmt, err := db.Prepare("UPDATE tasklog SET jInput  = ? WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Update the user's name and role.
	_, err = stmt.Exec(input, name)
	if err != nil {
		log.Fatal(err)
	}

	// Print a message to confirm the update.
	fmt.Println("Tasklog updated successfully!")

}

func handleFetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var task Task

	// Decode the JSON data from the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		// Handle the error, if any
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	db, err := sql.Open("mysql", "root:Ras$1203@tcp(localhost:3306)/editor")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// The values to be inserted
	name := task.Name

	fmt.Println("name:", name)
	query := `SELECT * FROM tasklog WHERE name = (?)`
	log.Println(query)
	rows, err3 := db.Query(query, name)
	if err3 != nil {
		panic(err3)
	}
	// defer rows.Close()
	// Name           string `json:"name"`
	// Type       string `json:"type"`
	// Parent         string `json:"parent"`
	// JInput         string `json:"jInput"`
	// Language    string `json:"language"`
	// JGroupAcess       []byte `json:"jGroupAcess"`
	// Flag bool   `json:"editpermission"`

	if true {
		for rows.Next() {
			fmt.Println("he haw")
			var taskcomp Taskcomp
			err := rows.Scan(&taskcomp.Name, &taskcomp.Type, &taskcomp.Parent, &taskcomp.JInput, &taskcomp.Language, &taskcomp.JGroupAcess, &taskcomp.Flag)
			if err != nil {
				log.Println("SQL ERROR", err)
			}
			var JGroupAcess []string
			err = json.Unmarshal(taskcomp.JGroupAcess, &JGroupAcess)
			if err != nil {
				log.Println("JSON", err)
			}
			// Inside the handleLogin function

			responseData := struct {
				Name        string `json:"name"`
				Type        string `json:"type"`
				Parent      string `json:"parent"`
				JInput      string `json:"jInput"`
				Language    string `json:"language"`
				JGroupAcess []byte `json:"jGroupAcess"`
				Flag        bool   `json:"editpermission"`
			}{
				Name:        taskcomp.Name,
				Type:        taskcomp.Type,
				Parent:      taskcomp.Parent,
				JInput:      taskcomp.JInput,
				Language:    taskcomp.Language,
				JGroupAcess: taskcomp.JGroupAcess,
				Flag:        taskcomp.Flag,
			}
			log.Println(responseData)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			w.WriteHeader(http.StatusOK)

		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	log.Println("POST HANDLE LOGIN")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var userlog Userlog

	// Decode the JSON data from the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&userlog)
	if err != nil {
		// Handle the error, if any
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	db, err := sql.Open("mysql", "root:Ras$1203@tcp(localhost:3306)/editor")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// The values to be inserted

	username := userlog.Username
	password := userlog.Password

	fmt.Println("tESTusername:", username)
	fmt.Println("tTETSpassword:", password)
	// The insert query with prepared statement
	query := `SELECT * FROM users WHERE username = (?)`
	// chk, err1 := db.Exec(query, email)
	// fmt.Println("chhh")
	// if err1 != nil {
	// 	panic(err1)
	// }
	// fmt.Println("chhk")
	// rowsAffected, err2 := chk.RowsAffected()
	// if err2 != nil {
	// 	panic(err2)
	// }
	// Execute the query
	log.Println(query)
	rows, err3 := db.Query(query, username)
	// fmt.Println("rowsAffected", rows.NextResultSet())

	if err3 != nil {
		panic(err3)
	}
	// defer rows.Close()

	if true {
		for rows.Next() {
			fmt.Println("he haw")
			var user User
			var id int64
			err := rows.Scan(&id, &user.Username, &user.Name, &user.Password, &user.Role, &user.Parent, &user.JGroup, &user.Permission, &user.Requests, &user.EditPermission)
			if err != nil {
				log.Println("SQL ERROR", err)
			}
			var jgroupArray []string
			err = json.Unmarshal(user.JGroup, &jgroupArray)
			if err != nil {
				log.Println("JSON", err)
			}
			// Inside the handleLogin function

			fmt.Println("user.username:", user.Username)
			fmt.Println("user.password:", user.Password)
			if password == user.Password {
				// ... other code ...
				responseData := struct {
					Username       string   `json:"username"`
					Name           string   `json:"name"`
					Password       string   `gorm:"password:password"`
					Role           string   `json:"role"`
					Parent         string   `json:"parent"`
					JGroup         []string `json:"jgroup"`
					Permission     string   `json:"permission"`
					Requests       string   `json:"requests"`
					EditPermission bool     `json:"editPermission"`
				}{
					Username:       user.Username,
					Name:           user.Name,
					Password:       user.Password,
					Role:           user.Role,
					Parent:         user.Parent,
					JGroup:         jgroupArray,
					Permission:     user.Permission,
					Requests:       user.Requests,
					EditPermission: user.EditPermission,
				}
				log.Println(responseData)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(responseData)
				w.WriteHeader(http.StatusOK)
			} else {
				// w.WriteHeader(una)
				w.WriteHeader(http.StatusUnauthorized)
			}

		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}

func main() {
	router := mux.NewRouter()

	// Define the route and handler for '/register/'
	router.HandleFunc("/register/", handleRegister).Methods("POST")
	router.HandleFunc("/login/", handleLogin).Methods("POST")
	router.HandleFunc("/commit/", handleCommit).Methods("POST")
	router.HandleFunc("/fetch/", handleFetch).Methods("POST")

	router.HandleFunc("/chk/", Chk).Methods("GET")
	router.HandleFunc("/task/", getTaks).Methods("GET")

	// Use CORS middleware
	corsMiddleware := cors.Default()
	// Wrap your router with the CORS middleware
	handler := corsMiddleware.Handler(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", handler))
}
