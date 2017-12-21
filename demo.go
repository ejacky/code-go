package main

import "github.com/gin-gonic/gin"
//import "github.com/gocolly/colly"

func serverLib() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}

func main() {
	serverLib();
}