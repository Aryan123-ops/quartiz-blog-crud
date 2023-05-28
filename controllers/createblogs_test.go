package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateBlog(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up the endpoint and handler function
	router.POST("/blog-post", CreateBlog)

	tests := []struct {
		name    string
		payload string
		status  int
		body    string
	}{
		{
			name:    "ValidPayload",
			payload: `{"title": "Test Title", "description": "Test Description", "body": "Test Body"}`,
			status:  http.StatusOK,
			body:    `{"data":{"title":"Test Title","description":"Test Description","body":"Test Body"}}`,
		},
		{
			name:    "MissingTitle",
			payload: `{"description": "Test Description", "body": "Test Body"}`,
			status:  http.StatusBadRequest,
			body:    `{"error":"Key: 'CreateBlogInput.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare a new HTTP request with the test payload
			req, err := http.NewRequest("POST", "/blog-post", strings.NewReader(tt.payload))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Create a new response recorder
			recorder := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(recorder, req)

			// Check the response status code
			assert.Equal(t, tt.status, recorder.Code)

			// Verify the response body
			assert.Equal(t, tt.body, recorder.Body.String())
		})
	}
}
