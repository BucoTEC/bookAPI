package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Book struct {
	ID       string    `json: "id"`
	Title    string `json: "title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book {
	{ID:"1", Title:"In search of lost time", Author:"Marcel Proust", Quantity: 2},
	{ID:"2", Title:"The great gatsby", Author:"Scot", Quantity: 5},
	{ID:"3", Title:"War and peac", Author:"Tolstoy", Quantity: 6},


}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context){
	var newBook Book

	if err := c.BindJSON(&newBook); err!= nil {
		return
	}

	books = append(books, newBook)
	c.Copy().IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks )
	router.Run("localhost:5000")
}