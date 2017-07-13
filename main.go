package main

import (
	"github.com/go-macaron/gzip"
	macaron "gopkg.in/macaron.v1"

	"github.com/mbier/aceleratomo/handlers"
)

func main() {

	m := macaron.New()
	m.Use(gzip.Gziper())

	m.Get("/", handlers.Raiz)
	m.Get("/quadro/geral", handlers.QuadroGeral)
	m.Get("/quadro/:nome", handlers.Quadro)

	m.Run()
}
