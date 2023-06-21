package handlers

import (
	"net/http"

	"github.com/Powerisinschool/image-x/converter"
	"github.com/Powerisinschool/image-x/loadx"
	"github.com/Powerisinschool/image-x/types"
	"github.com/gin-gonic/gin"
)

// ConvertFormatHandler handles the convert endpoint

//	@Summary		Convert Image Format
//	@Description	Convert the image format of an uploaded image
//	@Param			file	formData	file	true	"Image file to convert"
//	@Param			format	formData	string	true	"Target format to convert to"
//	@Success		200		{string}	string	"Conversion successful"
//	@Failure		400		{string}	string	"Bad request"
//	@Router			/convert [post]
func ConvertFormatHandler(c *gin.Context) {
	// Bind the request body to the ConvertFormatRequest struct
	var req types.ConvertFormatRequest
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

	buf, err := converter.Convert(img, req.Format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image", "message": err.Error()})
		return
	}

	// Return the converted image as the response
	c.Data(http.StatusOK, "image/"+req.Format, buf.Bytes())
}
