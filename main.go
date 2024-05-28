package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
    ID     uint   `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: 1, Title: "1984", Author: "George Orwell"},
    {ID: 2, Title: "Animal Farm", Author: "George Orwell"},
}

func main() {
    gin.SetMode(gin.ReleaseMode)

    router := gin.Default()
    router.SetTrustedProxies([]string{"192.168.1.100"}) // Sesuaikan dengan IP proxy Anda

    router.GET("/books", getBooks)
    router.GET("/books/:id", getBookByID)
    router.POST("/books", createBook)
    router.PUT("/books/:id", updateBook)
    router.DELETE("/books/:id", deleteBook)

    router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, book := range books {
        if fmt.Sprintf("%d", book.ID) == id {
            c.IndentedJSON(http.StatusOK, book)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func createBook(c *gin.Context) {
    var newBook Book
    if err := c.BindJSON(&newBook); err != nil {
        return
    }
    newBook.ID = uint(len(books) + 1)
    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book
    if err := c.BindJSON(&updatedBook); err != nil {
        return
    }
    for i, book := range books {
        if fmt.Sprintf("%d", book.ID) == id {
            books[i] = updatedBook
            updatedBook.ID = book.ID
            c.IndentedJSON(http.StatusOK, updatedBook)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBook(c *gin.Context) {
    id := c.Param("id")
    for i, book := range books {
        if fmt.Sprintf("%d", book.ID) == id {
            books = append(books[:i], books[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
