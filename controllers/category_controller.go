package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"sanbercode-mini-project/model"
	"sanbercode-mini-project/repository"

	"github.com/gin-gonic/gin"
)

type CreateCategoryInput struct {
	Name string `json:"name"`
}

func GetAllCategories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := repository.GetAllCategories(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if categories == nil {
			categories = []model.Category{}
		}
		c.JSON(http.StatusOK, categories)
	}
}

func CreateCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateCategoryInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category := model.Category{
			Name:       input.Name,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}

		err := repository.InsertCategory(db, &category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category created", "data": category})
	}
}

func GetCategoryByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		category, err := repository.GetCategoryByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusOK, category)
	}
}

func DeleteCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := repository.DeleteCategory(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found or failed to delete"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
	}
}

func UpdateCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var input CreateCategoryInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category := model.Category{
			Name:       input.Name,
			ModifiedAt: time.Now(),
		}

		err := repository.UpdateCategory(db, id, &category)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found or failed to update"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category updated", "data": category})
	}
}

func GetBooksByCategoryID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := repository.GetCategoryByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category ID not found"})
			return
		}
		
		books, err := repository.GetBooksByCategoryID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}