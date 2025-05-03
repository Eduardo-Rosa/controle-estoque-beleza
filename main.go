package main

import (
	"github.com/gin-gonic/gin"
	"controle-estoque-beleza/handler"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/products", handler.GetProducts)
	r.POST("/api/v1/products", handler.CreateProduct)
	r.PUT("/api/v1/products/:id", handler.UpdateProduct)
	r.DELETE("/api/v1/products/:id", handler.DeleteProduct)
	r.GET("/api/v1/products/search", handler.SearchProductsByCategory)

	r.Run(":8080")
}