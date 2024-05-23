package models

// Book represents a book object with its ID, title, and author.
type Book struct {
    ID     int    `json:"id"`      // ID uniquely identifies the book
    Title  string `json:"title"`   // Title is the title of the book
    Author string `json:"author"`  // Author is the author of the book
}
