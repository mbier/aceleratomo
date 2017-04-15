package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mbier/aceleratomo/handlers"
)

func main() {

	h := handlers.NewHandlerQuadro()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Raiz).Methods("GET")
	r.HandleFunc("/smofrete", h.QuadroTrack).Methods("GET")
	r.HandleFunc("/adm", h.QuadroAdm).Methods("GET")
	r.HandleFunc("/tms-web", h.QuadroTMSWEB).Methods("GET")
	r.HandleFunc("/smo-net", h.QuadroSMONET).Methods("GET")
	r.HandleFunc("/smo-web", h.QuadroSMOWEB).Methods("GET")
	r.HandleFunc("/smo-cte", h.QuadroSMOCTE).Methods("GET")
	r.HandleFunc("/delphi", h.QuadroDelhpi).Methods("GET")
	r.HandleFunc("/quadro-geral", handlers.QuadroGeral).Methods("GET")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
