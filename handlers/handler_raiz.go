package handlers

import (
	"fmt"
	"net/http"
)

// Raiz gerar html para /
func Raiz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<ul>"+
		"<li><a href=\"/quadro-geral\">Quadro Geral</li>"+
		"<li><a href=\"/smofrete\">SMOFRETE</li>"+
		"<li><a href=\"/adm\">Adm</li>"+
		"<li><a href=\"/tms-web\">TMS-WEB</li>"+
		"<li><a href=\"/smo-net\">SMO-NET</li>"+
		"<li><a href=\"/smo-web\">SMO-WEB</li>"+
		"<li><a href=\"/smo-cte\">SMO-CTE</li>"+
		"<li><a href=\"/delphi\">DELPHI</li>"+
		"</ul>")
}
