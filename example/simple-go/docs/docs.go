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
        "contact": {
            "name": "ZiGa",
            "email": "boringmanman@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users/v1/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Service"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseUsers"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Service"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "description": "主键ID",
                        "name": "id",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "邮箱",
                        "name": "email",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseNil"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/v1/:id/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Service"
                ],
                "summary": "用户下载",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "some id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseNil"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/v1/common/nil/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Service"
                ],
                "summary": "空参数示例1",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.User"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/v1/empty/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Service"
                ],
                "summary": "空参数示例2",
                "parameters": [
                    {
                        "type": "string",
                        "description": "参数无注释",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ResponseNil": {
            "type": "object"
        },
        "api.ResponseUsers": {
            "type": "object",
            "properties": {
                "data_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.User"
                    }
                }
            }
        },
        "api.User": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ShuaiGao.github.io",
	BasePath:         "/api/xxx",
	Schemes:          []string{},
	Title:            "xxx系统api文档",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}