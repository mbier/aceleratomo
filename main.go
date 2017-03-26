package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/mbier/aceleratomo/handlers"
	"os"
)

func main() {

	//mongoSession := mongo.GetMongoSession()

	//h := handlers.NewHandlerQuadro(mongoSession)
	h := handlers.NewHandlerQuadro(nil)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Raiz).Methods("GET")
	r.HandleFunc("/track", h.QuadroTrack).Methods("GET")
	r.HandleFunc("/adm", h.QuadroAdm).Methods("GET")
	r.HandleFunc("/tms-web", h.QuadroTMSWEB).Methods("GET")
	r.HandleFunc("/smo-net", h.QuadroSMONET).Methods("GET")
	r.HandleFunc("/smo-web", h.QuadroSMOWEB).Methods("GET")
	r.HandleFunc("/smo-cte", h.QuadroSMOCTE).Methods("GET")
	r.HandleFunc("/quadro-geral", handlers.QuadroGeral).Methods("GET")
	http.Handle("/", r)

	var port string = os.Getenv("PORT")
	if (port == nil) {
		port = "6969"
	}

	log.Fatal(http.ListenAndServe(":" + port, r))
}




