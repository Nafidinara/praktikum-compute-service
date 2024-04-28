package main

import (
	"Praktikum/database"
	"Praktikum/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.Migrate()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
