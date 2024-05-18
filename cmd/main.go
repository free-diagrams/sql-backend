package main

import (
	"github.com/free-diagrams/sql-backend/internal/app"

	_ "github.com/free-diagrams/sql-backend/docs"
)

// @title Free-Diagrams SQL API
// @version 1.0
// @description Free-Diagrams SQL API

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}
