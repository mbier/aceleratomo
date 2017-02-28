package main

import (
	"github.com/mbier/aceleratomo/acelerato"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/track", quadroTrackHandler).Methods("GET")
	r.HandleFunc("/tms-web", quadroTMSWEBHandler).Methods("GET")
	r.HandleFunc("/smo-net", quadroSMONETHandler).Methods("GET")
	r.HandleFunc("/smo-web", quadroSMOWEBHandler).Methods("GET")
	r.HandleFunc("/quadro-geral", quadroGeralHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":6969", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "/track \n" +
		"/tms-web \n" +
		"/smo-net \n" +
		"/smo-web \n")
}

func quadroTrackHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTrack())
}

func quadroTMSWEBHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTMSWEB())
}

func quadroSMONETHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMONET())
}

func quadroSMOWEBHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMOWEB())
}

func quadroGeralHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroGeral())
}
