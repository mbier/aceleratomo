package handlers

import (
	"strings"

	"net/http"

	"github.com/labstack/echo"
	"github.com/mbier/aceleratomo/projeto"
)

// Raiz gerar html para /
func Raiz(c echo.Context) error {
	html := "<ul>" +
		"<li><a href=\"/quadro/geral\">Quadro Geral</li>" +
		"<li><a href=\"/quadro/testes\">Quadro Testes</li>"

	for _, projeto := range projeto.GetProjetos() {
		html += "<li><a href=\"/quadro/" + strings.ToLower(projeto.Nome) + "\">" + projeto.Nome + "</li>"
	}

	html += "</ul>"

	return c.HTML(http.StatusOK, html)
}
