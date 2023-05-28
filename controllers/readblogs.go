package controllers

import (
	"blog-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBlog function implements functionality related  to fetch resource form the database.

// @Summary Get all blog posts
// @Description Get all added blog posts
// @Produce json
// @Success 200 {object} []models.Blog
// @Router /api/blog-post [get]
func GetBlog(c *gin.Context) {
	var blog []models.Blog
	models.DB.Find(&blog)

	c.JSON(http.StatusOK, map[string]interface{}{"data": blog})
}

// GetBlogByID function implements functionality related  to fetch resource form the database by id .

// @Summary Get a single blog post
// @Description Get a single blog post by ID
// @Produce json
// @Param id path int true "Blog Post ID"
// @Success 200 {object} models.Blog
// @Failure 400 {object} ErrorResponse
// @Router /api/blog-post/{id} [get]
func GetBlogByID(c *gin.Context) {
	var blog []models.Blog

	if err := models.DB.Where("id = ?", c.Param("id")).Find(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": blog})
}
