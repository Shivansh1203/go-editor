package routes

import (
	"github.com/gorilla/mux"
	"github.com/hrs24/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	// router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
