package main

import (
	"blog-crud/controllers"
	"blog-crud/models"
	"net/http"

	"github.com/swaggo/swag/example/basic/docs"

	"github.com/gin-gonic/gin"
)

func main() {

	docs.SwaggerInfo.Title = "blog-crud api"
	docs.SwaggerInfo.Description = "API documentation for the CRUD blog app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"

	//Defining routers as a default gin Engine.
	routers := gin.Default()

	// Calling Connect function that connects with Database.
	models.Connect()
	// Grouping a new router group by /api
	publicRoutes := routers.Group("/api")
	// publicRoutes.Static("/swagger", "./static")

	// Calling all the endpoints that will serve our database from the backend.
	publicRoutes.GET("/blog-post", controllers.GetBlog)
	publicRoutes.POST("/blog-post", controllers.CreateBlog)
	publicRoutes.GET("/blog-post/:id", controllers.GetBlogByID)
	publicRoutes.DELETE("/blog-post/:id", controllers.DeleteBlogById)
	publicRoutes.PATCH("/blog-post/:id", controllers.UpdateBlog)
	// Serve static files from the "static" folder
	//serve swagger documentation
	publicRoutes.GET("/swagger/*filepath", serveSwaggerUI)

	// Finally running all the endpoints on the "http://localhost:8080"
	routers.Run("localhost:8080")
}

func serveSwaggerUI(c *gin.Context) {
	http.FileServer(http.Dir("swagger")).ServeHTTP(c.Writer, c.Request)
}
