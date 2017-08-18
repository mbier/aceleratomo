package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mbier/aceleratomo/acelerato"
	"github.com/mbier/aceleratomo/projeto"
)

// APIProjeto gera uma quadro do projeto passado por parametro
func APIProjeto(c echo.Context) error {

	projeto, err := projeto.GetProjeto(c.Param("nome"))

	if err != nil {
		return err
	}
	quadro := acelerato.GerarDadosQuadro(projeto)

	return c.JSON(http.StatusOK, quadro)
}

// APIProjetos gera uma quadro do projeto passado por parametro
func APIProjetos(c echo.Context) error {

	quadros := acelerato.GerarDadosQuadros()

	return c.JSON(http.StatusOK, quadros)
}

// APIQuadroTeste gera uma quadro para testes
func APIQuadroTeste(c echo.Context) error {

	quadros := acelerato.GerarDadosQuadrosTeste()

	return c.JSON(http.StatusOK, quadros)
}
