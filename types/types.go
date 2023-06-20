package types

import "mime/multipart"

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

// ApplyFilterRequest represents the request body for the convert format endpoint
type ApplyFilterRequest struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`
	Filter string                `form:"filter" binding:"required,oneof=grayscale sepia blur sharpen"`
}
