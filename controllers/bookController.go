package controllers

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "bookapi/models"
)

// books represents a collection of Book objects
var books = []models.Book{
    {ID: 1, Title: "Book 1", Author: "Author 1"},
    {ID: 2, Title: "Book 2", Author: "Author 2"},
    {ID: 3, Title: "Book 3", Author: "Author 3"},
}

// nextID is used to generate the ID for the next book
var nextId = 4

// GetBooks returns the list of all books
func GetBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

// GetBookById returns a specific book based on the provided ID
func GetBookById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    book, err := findBookById(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
    var newBook models.Book
    if err := c.BindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newBook.ID = nextId
    nextId++
    books = append(books, newBook)
    c.JSON(http.StatusCreated, newBook)
}

// UpdateBook updates an existing book
func UpdateBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    var updatedBook models.Book
    if err := c.BindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, book := range books {
        if book.ID == id {
            // Check if title and author are provided in the request body,
            // if not, keep the existing values
            if updatedBook.Title != "" {
                books[i].Title = updatedBook.Title
            }
            if updatedBook.Author != "" {
                books[i].Author = updatedBook.Author
            }

            c.JSON(http.StatusOK, books) // Return the entire updated slice
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// DeleteBook deletes a book based on the provided ID
func DeleteBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    err = deleteBookByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// findBookById searches for a book by its ID
func findBookById(id int) (models.Book, error) {
    for _, book := range books {
        if book.ID == id {
            return book, nil
        }
    }
    return models.Book{}, errors.New("book not found")
}

// deleteBookByID deletes a book by its ID
func deleteBookByID(id int) error {
    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            return nil
        }
    }
    return errors.New("book not found")
}
