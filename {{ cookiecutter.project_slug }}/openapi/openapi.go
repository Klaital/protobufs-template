package openapi

import "embed"

//go:embed protos/*.swagger.json
var SwaggerDocs embed.FS

