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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT License",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/media": {
            "get": {
                "description": "Retrieve multiple media files by their IDs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Get multiple media files",
                "parameters": [
                    {
                        "description": "Array of Media IDs",
                        "name": "ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_media.Media"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Upload a new media file (image only)",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Upload a media file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_media.Media"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/media/{media_id}": {
            "get": {
                "description": "Retrieve a media file by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "media"
                ],
                "summary": "Get a media file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media ID",
                        "name": "media_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_media.Media"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a media file by its ID",
                "tags": [
                    "media"
                ],
                "summary": "Delete a media file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Media ID",
                        "name": "media_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "description": "Get all posts by a specific user address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get posts of user who made the request",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new post with title and content",
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
                        "description": "New Post",
                        "name": "newPost",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.NewPost"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/posts/all": {
            "get": {
                "description": "Get all posts with pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get all posts",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/posts/id/{id}": {
            "get": {
                "description": "Get a specific post by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get a post by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/posts/{address}": {
            "get": {
                "description": "Get all posts by a specific user address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get posts by user address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/posts/{address}/{index}": {
            "get": {
                "description": "Get a specific post by user address and index",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get a post by user address and index",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Post Index",
                        "name": "index",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login a user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "User Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_user.UserCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_user.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register a new user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "New User",
                        "name": "newUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_user.NewUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_user.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_St3plox_Blogchain_business_web_v1.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "big.Int": {
            "type": "object"
        },
        "github_com_St3plox_Blogchain_business_core_media.Media": {
            "type": "object",
            "properties": {
                "fileBytes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                }
            }
        },
        "github_com_St3plox_Blogchain_business_core_post.Category": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "Blog",
                "News",
                "Article"
            ]
        },
        "github_com_St3plox_Blogchain_business_core_post.NewPost": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Category"
                },
                "content": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_St3plox_Blogchain_business_core_post.Post": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "category": {
                    "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_post.Category"
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "$ref": "#/definitions/big.Int"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "timestamp": {
                    "$ref": "#/definitions/big.Int"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_St3plox_Blogchain_business_core_user.NewUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "github_com_St3plox_Blogchain_business_core_user.Role": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "github_com_St3plox_Blogchain_business_core_user.User": {
            "type": "object",
            "properties": {
                "address_hex": {
                    "type": "string"
                },
                "date_created": {
                    "type": "string"
                },
                "date_updated": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "private_key": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "public_key": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_St3plox_Blogchain_business_core_user.Role"
                    }
                },
                "secret_key": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "github_com_St3plox_Blogchain_business_core_user.UserCredentials": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "github_com_St3plox_Blogchain_business_web_v1.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "fields": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "jwtToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Blogchain API",
	Description:      "This is a sample server Blogchain server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
