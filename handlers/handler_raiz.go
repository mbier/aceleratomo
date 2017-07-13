package handlers

import (
	"strings"

	"github.com/mbier/aceleratomo/projeto"
	macaron "gopkg.in/macaron.v1"
)

// Raiz gerar html para /
func Raiz(ctx *macaron.Context) {
	html := "<ul>" +
		"<li><a href=\"/quadro/geral\">Quadro Geral</li>"

	for _, projeto := range projeto.GetProjetos() {
		html += "<li><a href=\"/quadro/" + strings.ToLower(projeto.Nome) + "\">" + projeto.Nome + "</li>"
	}

	html += "</ul>"

	ctx.Resp.Header().Set("Content-Type", "text/html")
	ctx.Resp.Write([]byte(html))
}
