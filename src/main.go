package main

import (
	"os"

	"example.com/redis/config/cache"
	"example.com/redis/config/database"
	"example.com/redis/features/item"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Init()
	cache.Init()

	r.GET("/", func(c *gin.Context) {
		service := &item.ItemService{}
		c.JSON(200, service.Get())
	})

	r.Run(":" + os.Getenv("PORTS"))
}
