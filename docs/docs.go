// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/convert": {
            "post": {
                "description": "Convert the image format of an uploaded image",
                "summary": "Convert Image Format",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file to convert",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Target format to convert to",
                        "name": "format",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Conversion successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/filter": {
            "post": {
                "description": "Apply the specified filter to the image",
                "consumes": [
                    "multipart/form-data"
                ],
                "summary": "Apply filter to image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file to apply filter",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter to apply: grayscale, sepia, blur, sharpen",
                        "name": "filter",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Filtered image",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request or parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/resize": {
            "post": {
                "description": "Resize the image",
                "consumes": [
                    "multipart/form-data"
                ],
                "summary": "Resize Image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file to resize",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Target width for resizing",
                        "name": "width",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Target height for resizing",
                        "name": "height",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Quality level for resizing: low, medium, high, best",
                        "name": "quality",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Resize successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Page Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Image Conversion API",
	Description:      "API endpoints for image conversion",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
