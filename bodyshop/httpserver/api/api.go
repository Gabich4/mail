package api

import (
	_ "bodyshop/docs" // этот пакет создается командой swag init -g internal/http/api.go
	_ "bodyshop/logic"
)

// @title Bodyshop API
// @version 1.0
// @description This is a sample articles server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1
// @query.collection.format multi

// @x-extension-openapi {"example": "value on a json format"}

