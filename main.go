package main

import (
	"os"

	"github.com/go-macaron/gzip"
	macaron "gopkg.in/macaron.v1"

	"strconv"

	"github.com/mbier/aceleratomo/handlers"
)

func main() {

	m := macaron.New()
	m.Use(gzip.Gziper())

	m.Get("/", handlers.Raiz)
	m.Get("/quadro/geral", handlers.QuadroGeral)
	m.Get("/quadro/:nome", handlers.Quadro)

	var port int
	portString := os.Getenv("PORT")

	if portString != "" {
		var err error
		port, err = strconv.Atoi(portString)
		panic(err)
	} else {
		port = 6969
	}

	m.Run(port)
}
