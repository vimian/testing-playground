package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vimian/testing-playground/pkg/types"
	"github.com/vimian/testing-playground/services/api_1/array"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, types.HTTPResponse{Status: "ok"})
	})

	r.POST("/min", func(c *gin.Context) {
		numbers := types.Float64s{}
		if err := c.BindJSON(&numbers); err != nil {
			c.JSON(400, types.HTTPResponse{Error: "Invalid input, expected a JSON array of numbers"})
			return
		}

		min, err := array.FindMin(numbers)
		if err != nil {
			c.JSON(400, types.HTTPResponse{Error: err.Error()})
			return
		}
		c.JSON(200, types.HTTPResponse{Result: min})
	})

	r.Run(":8080")
}
