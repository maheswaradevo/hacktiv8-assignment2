// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/license/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "get": {
                "description": "View all orders and return a JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "View all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.OrderResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new order from request body as a JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a New Order",
                "parameters": [
                    {
                        "description": "Create order",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.OrderResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateOrderRequest": {
            "type": "object",
            "properties": {
                "ItemsRequest": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemRequest"
                    }
                },
                "customer_name": {
                    "type": "string"
                }
            }
        },
        "dto.ItemRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "item_code": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "dto.ItemsResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "item_code": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer"
                },
                "order_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "dto.OrderResponse": {
            "type": "object",
            "properties": {
                "AllItems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemsResponse"
                    }
                },
                "customer_name": {
                    "type": "string"
                },
                "order_id": {
                    "type": "integer"
                },
                "ordered_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Orders API",
	Description:      "This is service to managing order",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
