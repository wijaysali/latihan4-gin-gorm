// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/customers/{id}": {
            "get": {
                "description": "Get Customer details by ID including address and orders",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieve customer",
                "operationId": "get-customer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Customer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Address": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "customerID": {
                    "type": "integer"
                }
            }
        },
        "model.Customer": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/model.Address"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Order"
                    }
                }
            }
        },
        "model.Order": {
            "type": "object",
            "properties": {
                "customerID": {
                    "type": "integer"
                },
                "orderDate": {
                    "type": "string"
                },
                "orderDetails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderDetail"
                    }
                },
                "totalPrice": {
                    "type": "number"
                }
            }
        },
        "model.OrderDetail": {
            "type": "object",
            "properties": {
                "orderID": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "productID": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
