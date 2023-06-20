package handlers

import (
	"net/http"

	"github.com/Powerisinschool/image-x/filter"
	"github.com/Powerisinschool/image-x/loadx"
	"github.com/Powerisinschool/image-x/types"
	"github.com/gin-gonic/gin"
)

// ApplyFilterHandler handles the apply filter endpoint

//	@Summary		Apply filter to image
//	@Description	Apply the specified filter to the image
//	@Accept			multipart/form-data
//	@Param			file	formData	file	true	"Image file to apply filter"
//	@Param			filter	formData	string	true	"Filter to apply: grayscale, sepia, blur, sharpen"
//	@Success		200		{string}	string	"Filtered image"
//	@Failure		400		{string}	string	"Invalid request or parameters"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/filter [post]
func ApplyFilterHandler(c *gin.Context) {
	// Handle the apply filter endpoint
	// TODO: Implement the logic to apply the specified filter to the image

	var req types.ApplyFilterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Access the uploaded file
	file, err := req.File.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer file.Close()

	// Implement the logic to convert the image format
	img, err := loadx.Load(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode image", "message": err.Error()})
		return
	}

	buf, err := filter.Filter(img, req.Filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image", "message": err.Error()})
		return
	}

	// Return the converted image as the response
	c.Data(http.StatusOK, "image/webp", buf.Bytes())
}
