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
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadro())
}
