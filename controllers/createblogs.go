package controllers

import (
	// "blog-crud/database"
	"blog-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Structure to create Blog Body and used in POST API.
type CreateBlogInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Body        string `json:"body"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

// @Summary Add a new blog post
// @Description Add a new blog post
// @Accept json
// @Produce json
// @Param post body CreateBlogInput true "Blog Post object"
// @Success 200 {object} CreateBlogInput
// @Failure 400 {object} ErrorResponse
// @Router /api/blog-post [post]
func CreateBlog(c *gin.Context) {
	// Validating the inputs
	var input CreateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	// Creating blog here
	book := models.Blog{Title: input.Title, Description: input.Description, Body: input.Body}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, map[string]interface{}{"data": book})
}
