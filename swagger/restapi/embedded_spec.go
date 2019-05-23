// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the API of the MedCo connector, that orchestrates the query at the MedCo node. It implements the PIC-SURE 2 API.",
    "title": "MedCo Connector",
    "contact": {
      "email": "medco-dev@listes.epfl.ch"
    },
    "license": {
      "name": "EULA",
      "url": "https://raw.githubusercontent.com/lca1/medco-connector/master/LICENSE"
    },
    "version": "0.2.0"
  },
  "basePath": "/medco-connector",
  "paths": {
    "/picsure2/info": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Returns information on how to interact with this PIC-SURE endpoint.",
        "operationId": "getInfo",
        "parameters": [
          {
            "description": "Credentials to be used.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/resourceInfo"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/picsure2/query": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Query MedCo node.",
        "operationId": "query",
        "parameters": [
          {
            "description": "Query request.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "query": {
                  "$ref": "#/definitions/query"
                },
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                },
                "resourceUUID": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/queryStatus"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/picsure2/query/sync": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Query MedCo node synchronously.",
        "operationId": "querySync",
        "parameters": [
          {
            "description": "Query request.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "query": {
                  "$ref": "#/definitions/query"
                },
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                },
                "resourceUUID": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/queryResult"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/picsure2/search": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Search through the ontology.",
        "operationId": "search",
        "parameters": [
          {
            "description": "Search request.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "query": {
                  "$ref": "#/definitions/searchQuery"
                },
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                },
                "resourceUUID": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/searchResult"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/picsure2/{queryId}/result": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Get result of query.",
        "operationId": "queryResult",
        "parameters": [
          {
            "type": "string",
            "description": "Query ID",
            "name": "queryId",
            "in": "path",
            "required": true
          },
          {
            "description": "Credentials to be used.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/queryResult"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/picsure2/{queryId}/status": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Get status of query.",
        "operationId": "queryStatus",
        "parameters": [
          {
            "type": "string",
            "description": "Query ID",
            "name": "queryId",
            "in": "path",
            "required": true
          },
          {
            "description": "Credentials to be used.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/queryStatus"
          },
          "default": {
            "$ref": "#/responses/error"
          }
        }
      }
    }
  },
  "definitions": {
    "query": {
      "type": "object",
      "properties": {
        "genomic-annotations": {
          "description": "genomic annotations query (todo)",
          "type": "object"
        },
        "i2b2-medco": {
          "description": "i2b2-medco query",
          "type": "object",
          "properties": {
            "differentialPrivacy": {
              "description": "differential privacy query parameters (todo)",
              "type": "object",
              "properties": {
                "queryBudget": {
                  "type": "number"
                }
              }
            },
            "panels": {
              "description": "i2b2 panels (linked by an AND)",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "items": {
                    "description": "i2b2 items (linked by an OR)",
                    "type": "array",
                    "items": {
                      "type": "object",
                      "properties": {
                        "encrypted": {
                          "type": "boolean"
                        },
                        "operator": {
                          "type": "string",
                          "enum": [
                            "exists",
                            "equals"
                          ]
                        },
                        "queryTerm": {
                          "type": "string"
                        },
                        "value": {
                          "type": "string"
                        }
                      }
                    }
                  },
                  "not": {
                    "description": "exclude the i2b2 panel",
                    "type": "boolean"
                  }
                }
              }
            },
            "queryType": {
              "$ref": "#/definitions/queryType"
            },
            "userPublicKey": {
              "type": "string"
            }
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "queryResultElement": {
      "type": "object",
      "properties": {
        "encryptedCount": {
          "type": "string"
        },
        "encryptedPatientList": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "encryptionKey": {
          "type": "string"
        },
        "queryType": {
          "$ref": "#/definitions/queryType"
        }
      }
    },
    "queryStatus": {
      "type": "object",
      "properties": {
        "duration": {
          "type": "integer",
          "format": "int64"
        },
        "expiration": {
          "type": "integer",
          "format": "int64"
        },
        "picsureResultId": {
          "type": "string"
        },
        "resourceID": {
          "type": "string"
        },
        "resourceResultId": {
          "type": "string"
        },
        "resourceStatus": {
          "type": "string"
        },
        "resultMetadata": {
          "type": "string",
          "format": "byte"
        },
        "sizeInBytes": {
          "type": "integer",
          "format": "int64"
        },
        "startTime": {
          "type": "integer",
          "format": "int64"
        },
        "status": {
          "type": "string",
          "enum": [
            "QUEUED",
            "PENDING",
            "ERROR",
            "AVAILABLE"
          ]
        }
      }
    },
    "queryType": {
      "type": "string",
      "enum": [
        "patient_list",
        "count_per_site",
        "count_per_site_obfuscated",
        "count_per_site_shuffled",
        "count_per_site_shuffled_obfuscated",
        "count_global",
        "count_global_obfuscated"
      ]
    },
    "resourceCredentials": {
      "type": "object",
      "properties": {
        "MEDCO_TOKEN": {
          "description": "MedCo authorization token.",
          "type": "string"
        }
      }
    },
    "searchQuery": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "children",
            "metadata"
          ]
        }
      }
    },
    "searchResultElement": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "leaf": {
          "type": "boolean"
        },
        "medcoEncryption": {
          "type": "object",
          "properties": {
            "childrenIds": {
              "type": "array",
              "items": {
                "type": "integer",
                "format": "int64"
              }
            },
            "encrypted": {
              "type": "boolean"
            },
            "id": {
              "type": "integer",
              "format": "int64"
            }
          }
        },
        "metadata": {
          "type": "object"
        },
        "name": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "container",
            "concept",
            "concept_numeric",
            "concept_enum",
            "concept_text",
            "genomic_annotation"
          ]
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "authorizations": {
          "type": "object",
          "properties": {
            "queryType": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/queryType"
              }
            }
          }
        },
        "id": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "error": {
      "description": "Error response",
      "schema": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    },
    "queryResult": {
      "description": "Query result.",
      "schema": {
        "$ref": "#/definitions/queryResultElement"
      }
    },
    "queryStatus": {
      "description": "Query status.",
      "schema": {
        "$ref": "#/definitions/queryStatus"
      }
    },
    "resourceInfo": {
      "description": "PIC-SURE 2 resource information.",
      "schema": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string"
          },
          "name": {
            "type": "string"
          },
          "queryFormats": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "description": {
                  "type": "string"
                },
                "examples": {
                  "type": "array",
                  "items": {
                    "type": "object"
                  }
                },
                "name": {
                  "type": "string"
                },
                "specifications": {
                  "type": "object"
                }
              }
            }
          }
        }
      }
    },
    "searchResult": {
      "description": "Search results.",
      "schema": {
        "type": "object",
        "properties": {
          "results": {
            "type": "array",
            "items": {
              "$ref": "#/definitions/searchResultElement"
            }
          },
          "searchQuery": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "MedCoToken": {
      "description": "MedCo JWT token.",
      "type": "oauth2",
      "flow": "application",
      "tokenUrl": "https://medco-demo.epfl.ch/auth"
    }
  },
  "security": [
    {
      "MedCoToken": null
    }
  ],
  "tags": [
    {
      "description": "PIC-SURE 2 Resource Service API",
      "name": "picsure2"
    }
  ],
  "externalDocs": {
    "description": "MedCo Technical Documentation",
    "url": "https://medco.epfl.ch/documentation"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the API of the MedCo connector, that orchestrates the query at the MedCo node. It implements the PIC-SURE 2 API.",
    "title": "MedCo Connector",
    "contact": {
      "email": "medco-dev@listes.epfl.ch"
    },
    "license": {
      "name": "EULA",
      "url": "https://raw.githubusercontent.com/lca1/medco-connector/master/LICENSE"
    },
    "version": "0.2.0"
  },
  "basePath": "/medco-connector",
  "paths": {
    "/picsure2/info": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Returns information on how to interact with this PIC-SURE endpoint.",
        "operationId": "getInfo",
        "parameters": [
          {
            "description": "Credentials to be used.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "PIC-SURE 2 resource information.",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "queryFormats": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "description": {
                        "type": "string"
                      },
                      "examples": {
                        "type": "array",
                        "items": {
                          "type": "object"
                        }
                      },
                      "name": {
                        "type": "string"
                      },
                      "specifications": {
                        "type": "object"
                      }
                    }
                  }
                }
              }
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/picsure2/query": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Query MedCo node.",
        "operationId": "query",
        "parameters": [
          {
            "description": "Query request.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "query": {
                  "$ref": "#/definitions/query"
                },
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                },
                "resourceUUID": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Query status.",
            "schema": {
              "$ref": "#/definitions/queryStatus"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/picsure2/query/sync": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Query MedCo node synchronously.",
        "operationId": "querySync",
        "parameters": [
          {
            "description": "Query request.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "query": {
                  "$ref": "#/definitions/query"
                },
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                },
                "resourceUUID": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Query result.",
            "schema": {
              "$ref": "#/definitions/queryResultElement"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/picsure2/search": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Search through the ontology.",
        "operationId": "search",
        "parameters": [
          {
            "description": "Search request.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "query": {
                  "$ref": "#/definitions/searchQuery"
                },
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                },
                "resourceUUID": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Search results.",
            "schema": {
              "type": "object",
              "properties": {
                "results": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/searchResultElement"
                  }
                },
                "searchQuery": {
                  "type": "string"
                }
              }
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/picsure2/{queryId}/result": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Get result of query.",
        "operationId": "queryResult",
        "parameters": [
          {
            "type": "string",
            "description": "Query ID",
            "name": "queryId",
            "in": "path",
            "required": true
          },
          {
            "description": "Credentials to be used.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Query result.",
            "schema": {
              "$ref": "#/definitions/queryResultElement"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/picsure2/{queryId}/status": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "picsure2"
        ],
        "summary": "Get status of query.",
        "operationId": "queryStatus",
        "parameters": [
          {
            "type": "string",
            "description": "Query ID",
            "name": "queryId",
            "in": "path",
            "required": true
          },
          {
            "description": "Credentials to be used.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "resourceCredentials": {
                  "$ref": "#/definitions/resourceCredentials"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Query status.",
            "schema": {
              "$ref": "#/definitions/queryStatus"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "query": {
      "type": "object",
      "properties": {
        "genomic-annotations": {
          "description": "genomic annotations query (todo)",
          "type": "object"
        },
        "i2b2-medco": {
          "description": "i2b2-medco query",
          "type": "object",
          "properties": {
            "differentialPrivacy": {
              "description": "differential privacy query parameters (todo)",
              "type": "object",
              "properties": {
                "queryBudget": {
                  "type": "number"
                }
              }
            },
            "panels": {
              "description": "i2b2 panels (linked by an AND)",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "items": {
                    "description": "i2b2 items (linked by an OR)",
                    "type": "array",
                    "items": {
                      "type": "object",
                      "properties": {
                        "encrypted": {
                          "type": "boolean"
                        },
                        "operator": {
                          "type": "string",
                          "enum": [
                            "exists",
                            "equals"
                          ]
                        },
                        "queryTerm": {
                          "type": "string"
                        },
                        "value": {
                          "type": "string"
                        }
                      }
                    }
                  },
                  "not": {
                    "description": "exclude the i2b2 panel",
                    "type": "boolean"
                  }
                }
              }
            },
            "queryType": {
              "$ref": "#/definitions/queryType"
            },
            "userPublicKey": {
              "type": "string"
            }
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "queryResultElement": {
      "type": "object",
      "properties": {
        "encryptedCount": {
          "type": "string"
        },
        "encryptedPatientList": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "encryptionKey": {
          "type": "string"
        },
        "queryType": {
          "$ref": "#/definitions/queryType"
        }
      }
    },
    "queryStatus": {
      "type": "object",
      "properties": {
        "duration": {
          "type": "integer",
          "format": "int64"
        },
        "expiration": {
          "type": "integer",
          "format": "int64"
        },
        "picsureResultId": {
          "type": "string"
        },
        "resourceID": {
          "type": "string"
        },
        "resourceResultId": {
          "type": "string"
        },
        "resourceStatus": {
          "type": "string"
        },
        "resultMetadata": {
          "type": "string",
          "format": "byte"
        },
        "sizeInBytes": {
          "type": "integer",
          "format": "int64"
        },
        "startTime": {
          "type": "integer",
          "format": "int64"
        },
        "status": {
          "type": "string",
          "enum": [
            "QUEUED",
            "PENDING",
            "ERROR",
            "AVAILABLE"
          ]
        }
      }
    },
    "queryType": {
      "type": "string",
      "enum": [
        "patient_list",
        "count_per_site",
        "count_per_site_obfuscated",
        "count_per_site_shuffled",
        "count_per_site_shuffled_obfuscated",
        "count_global",
        "count_global_obfuscated"
      ]
    },
    "resourceCredentials": {
      "type": "object",
      "properties": {
        "MEDCO_TOKEN": {
          "description": "MedCo authorization token.",
          "type": "string"
        }
      }
    },
    "searchQuery": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "children",
            "metadata"
          ]
        }
      }
    },
    "searchResultElement": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "leaf": {
          "type": "boolean"
        },
        "medcoEncryption": {
          "type": "object",
          "properties": {
            "childrenIds": {
              "type": "array",
              "items": {
                "type": "integer",
                "format": "int64"
              }
            },
            "encrypted": {
              "type": "boolean"
            },
            "id": {
              "type": "integer",
              "format": "int64"
            }
          }
        },
        "metadata": {
          "type": "object"
        },
        "name": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "enum": [
            "container",
            "concept",
            "concept_numeric",
            "concept_enum",
            "concept_text",
            "genomic_annotation"
          ]
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "authorizations": {
          "type": "object",
          "properties": {
            "queryType": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/queryType"
              }
            }
          }
        },
        "id": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "error": {
      "description": "Error response",
      "schema": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    },
    "queryResult": {
      "description": "Query result.",
      "schema": {
        "$ref": "#/definitions/queryResultElement"
      }
    },
    "queryStatus": {
      "description": "Query status.",
      "schema": {
        "$ref": "#/definitions/queryStatus"
      }
    },
    "resourceInfo": {
      "description": "PIC-SURE 2 resource information.",
      "schema": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string"
          },
          "name": {
            "type": "string"
          },
          "queryFormats": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "description": {
                  "type": "string"
                },
                "examples": {
                  "type": "array",
                  "items": {
                    "type": "object"
                  }
                },
                "name": {
                  "type": "string"
                },
                "specifications": {
                  "type": "object"
                }
              }
            }
          }
        }
      }
    },
    "searchResult": {
      "description": "Search results.",
      "schema": {
        "type": "object",
        "properties": {
          "results": {
            "type": "array",
            "items": {
              "$ref": "#/definitions/searchResultElement"
            }
          },
          "searchQuery": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "MedCoToken": {
      "description": "MedCo JWT token.",
      "type": "oauth2",
      "flow": "application",
      "tokenUrl": "https://medco-demo.epfl.ch/auth"
    }
  },
  "security": [
    {
      "MedCoToken": []
    }
  ],
  "tags": [
    {
      "description": "PIC-SURE 2 Resource Service API",
      "name": "picsure2"
    }
  ],
  "externalDocs": {
    "description": "MedCo Technical Documentation",
    "url": "https://medco.epfl.ch/documentation"
  }
}`))
}
