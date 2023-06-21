package main

import (
	"net/http"

	docs "github.com/Powerisinschool/image-x/docs"
	"github.com/Powerisinschool/image-x/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
	r.POST("/api/convert", handlers.ConvertFormatHandler)
	r.POST("/api/resize", handlers.ResizeHandler)
	r.POST("/api/filter", handlers.ApplyFilterHandler)
	r.POST("/api/upscale", handlers.UpscaleHandler)
	r.POST("/api/crop", cropImageHandler)
	r.POST("/api/rotate", rotateImageHandler)
	r.POST("/api/compress", compressImageHandler)
	r.POST("/api/thumbnail", generateThumbnailHandler)

	r.Run(":8080")
}

// Handler functions for your API endpoints

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
