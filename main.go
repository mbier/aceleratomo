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
	r.HandleFunc("/track", quadroTrackHandler).Methods("GET")
	r.HandleFunc("/tms-web", quadroTMSWEBHandler).Methods("GET")
	r.HandleFunc("/quadro-geral", quadroGeralHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func quadroTrackHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTrack())
}

func quadroTMSWEBHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTMSWEB())
}

func quadroGeralHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroGeral())
}
