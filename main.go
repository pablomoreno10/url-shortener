package main

import (
	"fmt"

	"github.com/pablomoreno10/url/handler"
	"github.com/pablomoreno10/url/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go URL Shortener API",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		handler.CreateShortUlr(c)
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		handler.HandlerShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the server - Error: %v", err))
	}
}
