// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Kelvin Gomes",
            "url": "https://www.linkedin.com/in/kelvin-gomes-fernandes",
            "email": "kelvingomesdeveloper@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/post": {
            "post": {
                "description": "Create a new post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create a new post",
                "parameters": [
                    {
                        "description": "Post Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePostInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePostOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorModel"
                        }
                    }
                }
            }
        },
        "/post/{id}": {
            "delete": {
                "description": "Delete post by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Delete post",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreatePostInput": {
            "type": "object",
            "properties": {
                "author_name": {
                    "type": "string"
                },
                "body": {
                    "type": "string"
                },
                "published_at": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePostOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "errors.ErrorModel": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {}
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Posts API",
	Description:      "Posts API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
