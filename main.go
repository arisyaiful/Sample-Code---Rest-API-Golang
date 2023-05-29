package main

import (
	"log"
	"sample-api/controllers"
	"sample-api/middleware"
	"sample-api/models"

	docs "sample-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(middleware.XssMiddleware())
	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.SecurityMiddleware())
	}
	r.Use(middleware.CorsMiddleware())

	models.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/helloworld", controllers.Helloworld)
		v1.GET("/books", controllers.FindBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.GET("/books/:id", controllers.FindBook)
		v1.PUT("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
		v1.GET("/category", controllers.FindCategory)
		v1.POST("/category", controllers.CreateCategory)
		v1.GET("/category/:id", controllers.FindCategoryById)
		v1.PUT("/category/:id", controllers.UpdateCategory)
		v1.DELETE("/category/:id", controllers.DeleteCategory)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
