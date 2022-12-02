package main

import (
	"os"

	"example.com/redis/config/cache"
	"example.com/redis/config/database"
	"example.com/redis/features/item"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message  string
	EnvValue string
}

func main() {
	r := gin.Default()
	database.Init()
	cache.Init()

	r.GET("/", func(c *gin.Context) {
		service := &item.ItemService{}
		c.JSON(200, service.Get())
	})

	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, Response{
			Message:  "success",
			EnvValue: os.Getenv("GIN_MODE"),
		})
	})

	r.Run(":9090")
}
