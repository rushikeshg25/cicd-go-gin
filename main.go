package main

import "github.com/gin-gonic/gin"

type Input struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/hello-post", func(ctx *gin.Context) {
		var input Input
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": input.Name,
		})
	})

	r.Run(":8081")
}
