package handlers

import (
	"net/http"

	"github.com/Powerisinschool/image-x/loadx"
	"github.com/Powerisinschool/image-x/resizer"
	"github.com/Powerisinschool/image-x/types"
	"github.com/gin-gonic/gin"
)

// ResizeHandler handles the resize endpoint

//	@Summary		Resize Image
//	@Description	Resize the image
//	@Accept			multipart/form-data
//	@Param			file	formData	file	true	"Image file to resize"
//	@Param			width	formData	int		true	"Target width for resizing"
//	@Param			height	formData	int		true	"Target height for resizing"
//	@Param			quality	formData	string	true	"Quality level for resizing: low, medium, high, best"
//	@Success		200		{string}	string	"Resize successful"
//	@Failure		400		{string}	string	"Bad request"
//	@Failure		404		{string}	string	"Page Not Found"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/resize [post]
func ResizeHandler(c *gin.Context) {
	// Bind the request body to the ConvertFormatRequest struct
	var req types.ResizeFormatRequest
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

	buf, err := resizer.Resize(img, req.Width, req.Height, req.Quality)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image", "message": err.Error()})
		return
	}

	// Return the converted image as the response
	c.Data(http.StatusOK, "image/webp", buf.Bytes())
}
