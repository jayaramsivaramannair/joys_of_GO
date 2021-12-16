package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/jayaramsivaramannair/joys_of_GO/Project_CRUD-API-With-GORM/pkg/utils"
	"github.com/jayaramsivaramannair/joys_of_GO/Project_CRUD-API-With-GORM/pkg/models"
)


var NewBook models.Book


func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json" )
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars.["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	//underscore is used since we do not want to use the db value returned from GetBookById function
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}