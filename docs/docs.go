// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "thaitanloi365@gmail.com"
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
        "/api/v1/admin/docker/container/list": {
            "get": {
                "description": "Get list docker container",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Docker"
                ],
                "summary": "Get list docker container",
                "responses": {}
            }
        },
        "/api/v1/admin/job_healthcheck": {
            "post": {
                "description": "Add healthcheck job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Job"
                ],
                "summary": "Add healthcheck job",
                "parameters": [
                    {
                        "description": "Form",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.HealthcheckJobCreateForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/admin/job_healthcheck/list": {
            "get": {
                "description": "Get list scheduled health check jobs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Job-HealthCheck"
                ],
                "summary": "Get list scheduled health check jobs",
                "responses": {}
            }
        },
        "/api/v1/admin/job_healthcheck/{tag}": {
            "delete": {
                "description": "Remove job by tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Job"
                ],
                "summary": "Remove job by tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag of job",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/admin/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Authorization"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        },
                        "headers": {
                            "Bearer": {
                                "type": "string",
                                "description": "OK"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/admin/me/logout": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Me"
                ],
                "summary": "Logout",
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/notifier/list": {
            "get": {
                "description": "Get list notifier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Notifier"
                ],
                "summary": "Get list notifier",
                "responses": {}
            }
        },
        "/api/v1/admin/notifier/{provider}": {
            "get": {
                "description": "Get notifier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Notifier"
                ],
                "summary": "Get notifier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag of job",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update notifier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Notifier"
                ],
                "summary": "Update notifier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag of job",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/docker/container/{id}": {
            "get": {
                "description": "Get list docker container",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Docker"
                ],
                "summary": "Get list docker container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of container",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/docker/container/{id}/stream_logs": {
            "get": {
                "description": "Stream docker container logs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Docker"
                ],
                "summary": "Stream docker container logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of container",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "errs.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "header": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.HealthcheckJobCreateForm": {
            "type": "object",
            "required": [
                "interval",
                "tag",
                "timeout"
            ],
            "properties": {
                "endpoint": {
                    "type": "string",
                    "example": "http://localhost:8080"
                },
                "interval": {
                    "type": "integer",
                    "example": 30
                },
                "tag": {
                    "type": "string",
                    "example": "api_healthcheck"
                },
                "timeout": {
                    "type": "integer",
                    "example": 20
                }
            }
        },
        "models.LoginForm": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_login": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "logged_out_at": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "timezone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "",
	Schemes:     []string{"http", "https"},
	Title:       "Go Monitor api docs",
	Description: "This is a go-monitor server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
