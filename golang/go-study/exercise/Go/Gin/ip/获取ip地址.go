package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/getip", func(c *gin.Context) {
		ip := c.ClientIP()
		c.JSON(http.StatusOK, gin.H{"ip": ip})
	})

	r.Run(":8080")
}
