package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Book struct {
	ID       string    `json:"id"`
	Title    string `json:"title"`
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
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context){
	id := c.Param("id")
	oneBook, err := getBookById(id)

	if err != nil{
		return
	}

	c.IndentedJSON(http.StatusOK,oneBook)

}

func getBookById(id string) (*Book, error ){
	for i,book := range books{
		if book.ID == id{
			return &books[i], nil
		}

	}
	return nil, errors.New("book not foun")
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks )
	router.GET("/books/:id", bookById )
	router.POST("/books", createBook )
	router.Run("localhost:5000")
}