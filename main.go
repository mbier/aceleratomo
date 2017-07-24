package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mbier/aceleratomo/handlers"
)

func main() {

	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.Raiz)
	e.GET("/quadro/geral", handlers.QuadroGeral)
	e.GET("/quadro/testes", handlers.QuadroTestes)
	e.GET("/quadro/:nome", handlers.Quadro)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
