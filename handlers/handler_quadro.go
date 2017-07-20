package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mbier/aceleratomo/acelerato"
	"github.com/mbier/aceleratomo/projeto"
)

// Quadro gera uma quadro do projeto passado por parametro
func Quadro(c echo.Context) error {

	projeto, err := projeto.GetProjeto(c.Param("nome"))

	if err != nil {
		return err
	}

	return c.HTML(http.StatusOK, acelerato.GerarQuadro(projeto))
}

// QuadroGeral gera um quadro com todos os projetos
func QuadroGeral(c echo.Context) error {
	return c.HTML(http.StatusOK, acelerato.GerarQuadroGeral())
}

// QuadroTestes gera um quadro especifico para testes com todos os projetos
func QuadroTestes(c echo.Context) error {
	return c.HTML(http.StatusOK, acelerato.GerarQuadroTestes())
}
