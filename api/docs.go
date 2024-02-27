// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/opening": {
            "get": {
                "description": "Find List Openings",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "List Openings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.OpeningListOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    }
                }
            },
            "put": {
                "description": "Update job opening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Update opening",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Opening id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.OpeningUpdateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.OpeningUpdateOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new job opening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Create opening",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.OpeningCreateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.OpeningCreateOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete job opening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Delete opening",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Opening id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.OpeningDeleteOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.OpeningResponse": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "service.ErrorOutput": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "string"
                }
            }
        },
        "service.OpeningCreateInput": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "service.OpeningCreateOutput": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "service.OpeningDeleteOutput": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "service.OpeningFindOutput": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "service.OpeningListOutput": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.OpeningResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "service.OpeningUpdateInput": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "service.OpeningUpdateOutput": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.OpeningResponse"
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
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
