package controllers

import (
	"net/http"

	"github.com/ashishsingh4u/bookmicroservice/models"
	"github.com/ashishsingh4u/bookmicroservice/repository"
	"github.com/gin-gonic/gin"
)

var repo repository.BookRepoInterface = &repository.BookRepository{}

// GET /books
// Get all books
func FindBooks(ctx *gin.Context) {
	var books []models.Book

	if err := repo.GetBooks(&books); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new book
func CreateBook(ctx *gin.Context) {
	// Validate input
	var input models.CreateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	var book models.Book
	if err := repo.CreateBook(&input, &book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a book
func FindBook(ctx *gin.Context) { // Get model if exist
	var book models.Book

	bookId := ctx.Param("id")
	if err := repo.GetBook(bookId, &book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(ctx *gin.Context) {
	// Validate input
	var input models.UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if code, err := repo.UpdateBook(ctx.Param("id"), &input, &book); err != nil {
		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(ctx *gin.Context) {

	if err := repo.DeleteBook(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "book deleted"})
}
