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
            "email": "support@devconnect.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users/login": {
            "post": {
                "description": "Authenticates a user with email and password, returning access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Login"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input or bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials or unverified email",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Failed to query user or generate tokens",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users Registration"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SignedUpUserOutput"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/users/reset-password": {
            "post": {
                "description": "Send a password reset token to the user's email address if they have a verified email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Password Reset Request"
                ],
                "summary": "sending Reset password request",
                "parameters": [
                    {
                        "description": "Password Reset Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.PasswordResetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset request success",
                        "schema": {
                            "$ref": "#/definitions/models.ResetRequestOutput"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid credentials or unverified email",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/users/update-password": {
            "post": {
                "description": "Send a password reset token to the user's email address if they have a verified email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reset Password"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "Password Reset Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdatePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset request success",
                        "schema": {
                            "$ref": "#/definitions/models.UpdatePasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid credentials or unverified email",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/users/verify-email": {
            "post": {
                "description": "Verifies a user's email by checking the verification token and updating the user's status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Verification"
                ],
                "summary": "Verify user email",
                "parameters": [
                    {
                        "description": "User email verification request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.EmailVerifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Email verified successfully",
                        "schema": {
                            "$ref": "#/definitions/models.VerifyEmailResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "401": {
                        "description": "Invalid or expired verification token",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "models.ImageInput": {
            "type": "object",
            "properties": {
                "base64String": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/models.UserResponse"
                }
            }
        },
        "models.ResetRequestOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.SignedUpUserOutput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "models.UpdatePasswordResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.VerifyEmailResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "requests.EmailVerifyRequest": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "object",
                    "required": [
                        "user_id",
                        "verification_token"
                    ],
                    "properties": {
                        "user_id": {
                            "type": "integer"
                        },
                        "verification_token": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "requests.LoginRequest": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "object",
                    "properties": {
                        "email": {
                            "type": "string"
                        },
                        "password": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "requests.PasswordResetRequest": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "object",
                    "required": [
                        "email"
                    ],
                    "properties": {
                        "email": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "requests.RegisterRequest": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "object",
                    "required": [
                        "email",
                        "password",
                        "phone",
                        "userName"
                    ],
                    "properties": {
                        "email": {
                            "type": "string"
                        },
                        "image": {
                            "$ref": "#/definitions/models.ImageInput"
                        },
                        "password": {
                            "type": "string",
                            "minLength": 6
                        },
                        "phone": {
                            "type": "string"
                        },
                        "role": {
                            "type": "string"
                        },
                        "userName": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "requests.UpdatePasswordRequest": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "object",
                    "required": [
                        "password",
                        "token",
                        "userId"
                    ],
                    "properties": {
                        "password": {
                            "type": "string",
                            "minLength": 6
                        },
                        "token": {
                            "type": "string"
                        },
                        "userId": {
                            "type": "integer"
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
	Host:             "localhost:5000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "User API fro cc tech project",
	Description:      "This is the API documentation for user api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
