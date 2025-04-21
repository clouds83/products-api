package main

import (
	"products-api/controller"
	"products-api/db"
	"products-api/repository"
	"products-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProducUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProducUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong", 
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.POST("/products", ProductController.CreateProduct)

	server.GET("/products/:productId", ProductController.GetProductById)

	server.Run(":8000")
}