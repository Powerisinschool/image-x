package handlers

import (
	"net/http"

	"github.com/Powerisinschool/image-x/loadx"
	"github.com/Powerisinschool/image-x/types"
	"github.com/Powerisinschool/image-x/upscaler"
	"github.com/gin-gonic/gin"
)

// UpscaleHandler handles the upscale endpoint

//	@Summary		Upscale image
//	@Description	Upscale the image by the specified scale
//	@Accept			multipart/form-data
//	@Param			file	formData	file	true	"Image file to upscale"
//	@Param			scale	query		int		false	"Upscale scale (default: 2)"
//	@Success		200		{string}	string	"Upscaled image"
//	@Failure		400		{string}	string	"Invalid request or parameters"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/upscale [post]
func UpscaleHandler(c *gin.Context) {
	// Handle the upscale endpoint

	var req types.UpscaleRequest
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

	// Implement the logic to upscale the image
	img, err := loadx.Load(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode image", "message": err.Error()})
		return
	}

	buf, err := upscaler.Upscale(img)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upscale image", "message": err.Error()})
		return
	}

	// Return the upscaled image as the response
	c.Data(http.StatusOK, "image/jpeg", buf.Bytes())
}
