package models

import (
	"github.com/hrs24/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email          string `gorm:""json:"email"`
	Name           string `"json:"name"`
	Password       string `"json:"password"`
	Role           string `"json:"role"`
	Parent         string `"json:"parent"`
	JGroup         []byte `json:"author"`
	Permission     string `"json:"permission"`
	Requests       string `"json:"requests"`
	EditPermission string `"json:"editpermission"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// func GetAllBooks() []Book {
// 	var Books []Book
// 	db.Find(&Books)
// 	return Books
// }

// func GetBookById(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	db := db.Where("ID=?", Id).Find(&getBook)
// 	return &getBook, db
// }

// func DeleteBook(ID int64) Book {
// 	var book Book
// 	db.Where("ID=?", ID).Delete(book)
// 	return book
// }
