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
        "/": {
            "get": {
                "tags": [
                    "App"
                ],
                "summary": "test serverless function",
                "operationId": "exec-app",
                "responses": {}
            }
        },
        "/api/projects": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get projects",
                "operationId": "get-projects",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "pageNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/project.GenericProjectResp"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Create project",
                "operationId": "create-project",
                "parameters": [
                    {
                        "description": "create project DTO",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/projects.CreateProjectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/project.GenericProjectResp"
                        }
                    }
                }
            }
        },
        "/api/projects/code/{projectId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Download project",
                "operationId": "download-project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Upload project",
                "operationId": "upload-project-code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "zipped code",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/projects/envs/{projectId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Env"
                ],
                "summary": "Get Environmental Variables",
                "operationId": "get-env-variables",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/project.EnvVariable"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Env"
                ],
                "summary": "Create Environmental Variables",
                "operationId": "create-env-variables",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create env variables",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/envs.OverwriteEnvVariablesReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/project.EnvVariable"
                            }
                        }
                    }
                }
            }
        },
        "/api/projects/logs/{projectId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "Get Project log",
                "operationId": "get-project-log",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "pageNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Log type",
                        "name": "logType",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "searchQuery",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "Delete Project logs",
                "operationId": "delete-project-log",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/plans": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Get Pricing Plans",
                "operationId": "get-pricing-plans",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "pageNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get project",
                "operationId": "get-project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/projects.ProjectApiResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Update project",
                "operationId": "update-project",
                "parameters": [
                    {
                        "description": "create project DTO",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/projects.UpdateProjectRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/project.GenericProjectResp"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Delete project",
                "operationId": "delete-project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/projects/{id}/api-token": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Generate API token",
                "operationId": "generate-api-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Remove API Key",
                "operationId": "remove-api-key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/{projectId}/analytics/hourly": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Analytics"
                ],
                "summary": "Get Hourly Project Invocations",
                "operationId": "get-project-hourly-invocations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Invocation Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/{projectId}/analytics/monthly": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Analytics"
                ],
                "summary": "Get Monthly Project Invocations",
                "operationId": "get-project-monthly-invocations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Invocation Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/{projectId}/analytics/weekly": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Analytics"
                ],
                "summary": "Get Weekly Project Invocations",
                "operationId": "get-project-weekly-invocations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Invocation Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/{projectId}/plan": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Set Project Pricing Plan",
                "operationId": "set-project-pricing-plan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Plan Id",
                        "name": "planId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/api/projects/{projectId}/resources": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Resources"
                ],
                "summary": "Get Project Resource Consumption",
                "operationId": "get-project-resource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "pageNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Month",
                        "name": "month",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/{projectId}/resources/monthly": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Resources"
                ],
                "summary": "Get Monthly Project Resource Consumption",
                "operationId": "get-monthly-project-resource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Month",
                        "name": "month",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/projects/{projectId}/resources/total": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Resources"
                ],
                "summary": "Get Total Project Resource Consumptions",
                "operationId": "get-total-project-resource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "Get users",
                "operationId": "get-users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "pageNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "Change user account status",
                "operationId": "change-users-account-status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/api/users/{id}/projects": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "Delete users projects",
                "operationId": "delete-users-projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/app/status": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "update serverless app status",
                "operationId": "app-status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "App subdomain",
                        "name": "subdomain",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "envs.EnvVariables": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "envs.OverwriteEnvVariablesReq": {
            "type": "object",
            "properties": {
                "variables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/envs.EnvVariables"
                    }
                }
            }
        },
        "project.DeploymentStatus": {
            "type": "string",
            "enum": [
                "failed",
                "building",
                "done",
                "none"
            ],
            "x-enum-varnames": [
                "Failed",
                "Building",
                "Done",
                "None"
            ]
        },
        "project.EnvVariable": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "projectID": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "project.GenericProjectResp": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deploymentStatus": {
                    "$ref": "#/definitions/project.DeploymentStatus"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "projects.CreateProjectRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "projects.ProjectApiResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deploymentStatus": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "projects.UpdateProjectRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pulsar API",
	Description:      "This is a server for  pulsar (serverless web platform) server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
