package controllers

import (
	"blog-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteBlog function implements functionality related  to delete resource in the databaseby id .

// @Summary Delete a blog post
// @Description Delete a blog post by ID
// @Produce json
// @Param id path int true "Blog Post ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Router /api/blog-post/{id} [delete]
func DeleteBlogById(c *gin.Context) {
	var blog models.Blog

	if err := models.DB.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "records not found"})
		return
	}
	models.DB.Delete(&blog)
	c.JSON(http.StatusOK, map[string]interface{}{"data": "Deleted Successfully"})
}
