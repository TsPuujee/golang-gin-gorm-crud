package main

import (
	"gorm-test/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	bookRepo := controllers.New()
	r.POST("/books", bookRepo.CreateBook)
	r.GET("/books", bookRepo.GetBooks)
	r.GET("/books/:id", bookRepo.GetBook)
	r.PUT("/books/:id", bookRepo.UpdateBook)
	r.DELETE("/books/:id", bookRepo.DeleteBook)

	return r
}
