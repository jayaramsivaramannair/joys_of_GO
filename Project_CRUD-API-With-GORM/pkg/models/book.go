package models

import (
	"gorm.io/gorm"
	"github.com/jayaramsivaramannair/joys_of_GO/Project_CRUD-API-With-GORM/pkg/config"
)

var db *gorm.DB


type Book struct {
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//Creates a new book in the database
func (b *Book) CreateBook() *Book {
	//db.NewRecord(b)
	db.Create(&b)
	return b
}

//Get all the books currently stored in the database
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}


//Get a specific book by its id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}


//Function to delete a specific book from the database
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

