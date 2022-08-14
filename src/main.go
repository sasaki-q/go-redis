package main

import (
	"os"

	"example.com/redis/config/cache"
	"example.com/redis/config/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Init()
	cache.Init()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World")
	})

	r.Run(":" + os.Getenv("PORTS"))
}
