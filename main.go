package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:quantity`
}

var books = []book{
	{ID: "1", Title: "前端三十", Author: "朱信穎", Quantity: 10},
	{ID: "2", Title: "精通Python", Author: "Bill Lubanovic", Quantity: 1},
}

// *gin.Context It carries request details, validates and serializes JSON, and more.
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context){
	var newBook book

	// 用c.BindJSON把body綁在變數上
	if err := c.BindJSON(&newBook);err != nil{
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook) // 序列化struct並且response
	// c.IndentedJSON vs built in context.JSON. 前者較簡短，debug容易
}

func bookById(c *gin.Context){
	id := c.Param("id") // find path param
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error){
	for _, book := range books{
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("沒有這本書")
}

func main(){
	// router
	r := gin.Default()
	r.GET("/books/:id", bookById)
	r.GET("/books", getBooks)
	r.POST("/books", createBook)
	r.Run() // attach the router to an http.Server and start the server
}