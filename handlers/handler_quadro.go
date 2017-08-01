package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mbier/aceleratomo/projeto"
	"github.com/mbier/aceleratomo/quadro"
)

// Quadro gera uma quadro do projeto passado por parametro
func Quadro(c echo.Context) error {

	projeto, err := projeto.GetProjeto(c.Param("nome"))

	if err != nil {
		return err
	}

	return c.HTML(http.StatusOK, quadro.GerarQuadro(projeto))
}

// QuadroGeral gera um quadro com todos os projetos
func QuadroGeral(c echo.Context) error {
	return c.HTML(http.StatusOK, quadro.GerarQuadroGeral())
}

// QuadroTestes gera um quadro especifico para testes com todos os projetos
func QuadroTestes(c echo.Context) error {
	return c.HTML(http.StatusOK, quadro.GerarQuadroTestes())
}
