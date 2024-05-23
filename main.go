package main

import (
    "github.com/gin-gonic/gin"      // Import Gin framework
    "bookapi/controllers"            // Import controllers package where route handlers are defined
)

func main() {
    r := gin.Default()  // Create a Gin router with default middleware

    // Define routes for CRUD operations on books
    r.GET("/books", controllers.GetBooks)             // GET /books - Retrieves a list of all books
    r.GET("/books/:id", controllers.GetBookById)      // GET /books/:id - Retrieves a specific book by ID
    r.POST("/books", controllers.CreateBook)          // POST /books - Creates a new book
    r.PUT("/books/:id", controllers.UpdateBook)       // PUT /books/:id - Updates an existing book by ID
    r.DELETE("/books/:id", controllers.DeleteBook)    // DELETE /books/:id - Deletes a book by ID

    r.Run(":8080")  // Start the Gin server and listen on port 8080
}
