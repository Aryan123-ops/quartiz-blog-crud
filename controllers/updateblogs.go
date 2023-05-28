package controllers

import (
	"blog-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Structure to update Blog Body and used in PATCH API.
type UpdateBlogInput struct {
	Title       string `json:"title"`
	Description string `json:"description" `
	Body        string `json:"body"`
}

// UpdateBlog function implements functionality related  to update resource in the databaseby id .

// @Summary Update a blog post
// @Description Update a blog post by ID
// @Accept json
// @Produce json
// @Param id path int true "Blog Post ID"
// @Param post body models.Blog true "Updated Blog Post object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Router /api/blog-post/{id} [patch]
func UpdateBlog(c *gin.Context) {
	var blog models.Blog

	if err := models.DB.Where("id  = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Record Not Found!"})
		return
	}

	var input UpdateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	models.DB.Model(&blog).Updates(input)
	c.JSON(http.StatusOK, map[string]interface{}{"data": blog})
}
