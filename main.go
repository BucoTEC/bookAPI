package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		c.IndentedJSON(http.StatusInternalServerError, err)
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
	return nil, errors.New("book not found")
}
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}





func main() {
	router := gin.Default()

	router.GET("/books", getBooks )
	router.GET("/books/:id", bookById )
	router.POST("/books", createBook )
	router.PATCH("/checkout", checkoutBook)
	router.PUT("/checkin", returnBook)

	router.Run("localhost:5000")
	dsn := "host=localhost user=postgres password=04012000 dbname=Books_API port=5432 sslmode=disable TimeZone=Europe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	
	
	fmt.Printf("Connectio do database is open : %v", db)
}
