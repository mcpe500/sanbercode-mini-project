package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sanbercode-mini-project/model"
	"sanbercode-mini-project/repository"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"total_page"`
	CategoryID  int    `json:"category_id"`
}

func CreateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if strings.TrimSpace(input.Title) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
			return
		}

		if input.Price <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than 0"})
			return
		}

		if input.TotalPage <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Total page must be greater than 0"})
			return
		}

		if input.ReleaseYear < 1980 || input.ReleaseYear > 2024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Release Year harus antara 1980 dan 2024"})
			return
		}

		var thickness string
		if input.TotalPage > 100 {
			thickness = "tebal"
		} else {
			thickness = "tipis"
		}

		book := model.Book{
			Title:       input.Title,
			Description: input.Description,
			ImageUrl:    input.ImageURL,
			ReleaseYear: input.ReleaseYear,
			Price:       input.Price,
			TotalPage:   input.TotalPage,
			Thickness:   thickness,
			CategoryID:  input.CategoryID,
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
		}

		err := repository.InsertBook(db, &book)
		if err != nil {
			if strings.Contains(err.Error(), "foreign key constraint") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id: category does not exist"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil ditambahkan", "data": book})
	}
}

func GetAllBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := repository.GetAllBooks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if books == nil {
			books = []model.Book{}
		}
		c.JSON(http.StatusOK, books)
	}
}

func GetBookByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		book, err := repository.GetBookByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := repository.DeleteBook(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	}
}

func UpdateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var input CreateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if strings.TrimSpace(input.Title) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
			return
		}
		if input.Price <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than 0"})
			return
		}

		if input.TotalPage <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Total page must be greater than 0"})
			return
		}

		if input.ReleaseYear < 1980 || input.ReleaseYear > 2024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Release Year harus antara 1980 dan 2024"})
			return
		}

		var thickness string
		if input.TotalPage > 100 {
			thickness = "tebal"
		} else {
			thickness = "tipis"
		}

		book := model.Book{
			Title:       input.Title,
			Description: input.Description,
			ImageUrl:    input.ImageURL,
			ReleaseYear: input.ReleaseYear,
			Price:       input.Price,
			TotalPage:   input.TotalPage,
			Thickness:   thickness,
			CategoryID:  input.CategoryID,
			ModifiedAt:  time.Now(),
		}

		err := repository.UpdateBook(db, id, &book)
		if err != nil {
			if strings.Contains(err.Error(), "foreign key constraint") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id: category does not exist"})
				return
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or failed to update"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil diupdate", "data": book})
	}
}