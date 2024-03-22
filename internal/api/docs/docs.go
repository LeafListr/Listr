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
        "/dispensaries": {
            "get": {
                "description": "Returns a list of supported dispensaries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dispensaries"
                ],
                "summary": "List supported dispensaries",
                "responses": {
                    "200": {
                        "description": "List of supported dispensaries\"\tEnums(Curaleaf, Beyond-Hello)",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}": {
            "get": {
                "description": "Returns details of a specific dispensary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dispensaries"
                ],
                "summary": "Get dispensary details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dispensary details",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.supportedDispensaryOptions"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations": {
            "get": {
                "description": "Returns a list of locations for a specific dispensary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "locations"
                ],
                "summary": "List locations for a dispensary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Recreational or medical",
                        "name": "recreational",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of locations",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Location"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations/{locationId}": {
            "get": {
                "description": "Returns details of a specific location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "locations"
                ],
                "summary": "Get location details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "locationId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Location details",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations/{locationId}/cannabinoids": {
            "get": {
                "description": "Returns a list of cannabinoids for a specific location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cannabinoids"
                ],
                "summary": "List cannabinoids for a location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "locationId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Recreational or medical",
                        "name": "recreational",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of cannabinoids",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cannabinoid"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations/{locationId}/categories": {
            "get": {
                "description": "Returns a list of categories for a specific location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "List categories for a location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "locationId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Recreational or medical",
                        "name": "recreational",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of categories",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations/{locationId}/offers": {
            "get": {
                "description": "Returns a list of offers for a specific location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "List offers for a location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "locationId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Recreational or medical",
                        "name": "recreational",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of offers",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Offer"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations/{locationId}/products": {
            "get": {
                "description": "Returns a list of products for a specific location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "List products for a location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "locationId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Recreational or medical",
                        "name": "recreational",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sub Category",
                        "name": "sub",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Minimum price per gram",
                        "name": "min_price_per_g",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Maximum price per gram",
                        "name": "max_price_per_g",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Minimum price",
                        "name": "min_price",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Maximum price",
                        "name": "max_price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Brands to include",
                        "name": "brands",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Brands to exclude",
                        "name": "not_brands",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Variants to include",
                        "name": "variants",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Terms to exclude",
                        "name": "excludes",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Terms to include",
                        "name": "includes",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "Sort products by price",
                        "name": "price_sort",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "Sort products by THC",
                        "name": "thc_sort",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "Sort products by Total Terpenes",
                        "name": "terp_sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Most important terpene",
                        "name": "terp1",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Second most important terpene",
                        "name": "terp2",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Third most important terpene",
                        "name": "terp3",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Sort terpene profile low to high",
                        "name": "terp_asc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of products",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations/{locationId}/products/{productId}": {
            "get": {
                "description": "Returns details of a specific product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "locationId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Recreational or medical",
                        "name": "recreational",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "productId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product details",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/dispensaries/{dispensaryId}/locations/{locationId}/terpenes": {
            "get": {
                "description": "Returns a list of terpenes for a specific location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "terpenes"
                ],
                "summary": "List terpenes for a location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dispensary ID",
                        "name": "dispensaryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "locationId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Recreational or medical",
                        "name": "recreational",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of terpenes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Terpene"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.supportedDispensaryOptions": {
            "type": "string",
            "enum": [
                "locations"
            ],
            "x-enum-varnames": [
                "Locations"
            ]
        },
        "models.Cannabinoid": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "models.Location": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "zipCode": {
                    "type": "string"
                }
            }
        },
        "models.Offer": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "models.Price": {
            "type": "object",
            "properties": {
                "discountedTotal": {
                    "type": "number"
                },
                "isDiscounted": {
                    "type": "boolean"
                },
                "perGram": {
                    "type": "number"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "cannabinoids": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Cannabinoid"
                    }
                },
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "permalink": {
                    "type": "string"
                },
                "price": {
                    "$ref": "#/definitions/models.Price"
                },
                "subcategory": {
                    "type": "string"
                },
                "terpeneTotal": {
                    "type": "number"
                },
                "terpenes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Terpene"
                    }
                },
                "variant": {
                    "type": "string"
                }
            }
        },
        "models.Terpene": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.2.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Listr API",
	Description:      "This is the Listr server for dispensary management.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
