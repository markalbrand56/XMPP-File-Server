// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/files/:directory/:file": {
            "get": {
                "description": "Obtiene un archivo de S3",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Obtiene un archivo de S3",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Directorio donde se encuentra el archivo",
                        "name": "directory",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Nombre del archivo",
                        "name": "file",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.StandardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.StandardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.StandardResponse"
                        }
                    }
                }
            }
        },
        "/files/{directory}": {
            "post": {
                "description": "(Requiere autentificación) Sube un archivo a S3 y lo asocia al usuario que lo subió. Necesita un token de autentificación",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "(Requiere autentificación) Sube un archivo a S3 y lo asocia al usuario que lo subió",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Directorio donde se guardará el archivo",
                        "name": "directory",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Archivo a subir",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UploadSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.StandardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.StandardResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.StandardResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "responses.UploadSuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "paths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
