package main

import (
	"mime/multipart"
	"net/http"

	"github.com/Powerisinschool/image-x/converter"
	docs "github.com/Powerisinschool/image-x/docs"
	"github.com/Powerisinschool/image-x/loadx"
	resizer "github.com/Powerisinschool/image-x/resizer"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ConvertFormatRequest represents the request body for the convert format endpoint
type ConvertFormatRequest struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`
	Format string                `form:"format" binding:"required"`
}

// ResizeFormatRequest represents the request body for the convert format endpoint
type ResizeFormatRequest struct {
	File    *multipart.FileHeader `form:"file" binding:"required"`
	Width   int                   `form:"width" binding:"required"`
	Height  int                   `form:"height" binding:"required"`
	Quality string                `form:"quality" binding:"required,oneof=low medium high best"`
}

//	@title			Image Conversion API
//	@version		1.0
//	@description	API endpoints for image conversion
//	@BasePath		/api
func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "ImageX API"
	docs.SwaggerInfo.Description = "This is an Image Conversion server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Define API documentation
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define the API routes and handlers
	r.POST("/api/convert", convertFormatHandler)
	r.POST("/api/resize", resizeHandler)
	r.POST("/api/filter", applyFilterHandler)
	r.POST("/api/crop", cropImageHandler)
	r.POST("/api/rotate", rotateImageHandler)
	r.POST("/api/compress", compressImageHandler)
	r.POST("/api/thumbnail", generateThumbnailHandler)

	r.Run(":8080")
}

// Handler functions for your API endpoints

//	@Summary		Convert Image Format
//	@Description	Convert the image format of an uploaded image
//	@Param			file	formData	file	true	"Image file to convert"
//	@Param			format	formData	string	true	"Target format to convert to"
//	@Success		200		{string}	string	"Conversion successful"
//	@Failure		400		{string}	string	"Bad request"
//	@Router			/convert [post]
func convertFormatHandler(c *gin.Context) {
	// Bind the request body to the ConvertFormatRequest struct
	var req ConvertFormatRequest
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

// resizeHandler handles the resize endpoint

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
func resizeHandler(c *gin.Context) {
	// Bind the request body to the ConvertFormatRequest struct
	var req ResizeFormatRequest
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

func applyFilterHandler(c *gin.Context) {
	// Handle the apply filter endpoint
	// TODO: Implement the logic to apply the specified filter to the image
}

func cropImageHandler(c *gin.Context) {
	// Handle the crop image endpoint
	// TODO: Implement the logic to crop the image based on the provided coordinates/dimensions
}

func rotateImageHandler(c *gin.Context) {
	// Handle the rotate image endpoint
	// TODO: Implement the logic to rotate the image by the specified angle
}

func compressImageHandler(c *gin.Context) {
	// Handle the compress image endpoint
	// TODO: Implement the logic to compress the image while preserving quality
}

func generateThumbnailHandler(c *gin.Context) {
	// Handle the generate thumbnail endpoint
	// TODO: Implement the logic to generate a thumbnail image of the specified size
}
