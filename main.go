package main

import(
	"fmt"
	"log"
	"os"

	"sanbercode-mini-project/controllers"
	"sanbercode-mini-project/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()

	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrations",
	}

	n, err := migrate.Exec(database.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Gagal migrasi: ", err)
	}
	fmt.Println("Migrasi berhasil: ", n)

	r:= gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	auth := gin.BasicAuth(gin.Accounts{
		"admin": "password123",
	})
	
	api := r.Group("/api")
	{
		api.GET("/books", auth, controllers.GetAllBooks(database.DB))
		api.GET("/books/:id", auth, controllers.GetBookByID(database.DB))
		api.POST("/books", auth, controllers.CreateBook(database.DB))
		api.DELETE("/books/:id", auth, controllers.DeleteBook(database.DB))
		api.PUT("/books/:id", auth, controllers.UpdateBook(database.DB))

		api.GET("/categories", auth, controllers.GetAllCategories(database.DB))
		api.GET("/categories/:id", auth, controllers.GetCategoryByID(database.DB))
		api.POST("/categories", auth, controllers.CreateCategory(database.DB))
		api.DELETE("/categories/:id", auth, controllers.DeleteCategory(database.DB))
		api.PUT("/categories/:id", auth, controllers.UpdateCategory(database.DB))
		api.GET("/categories/:id/books", auth, controllers.GetBooksByCategoryID(database.DB))
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}