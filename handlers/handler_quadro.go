package handlers

import (
	"fmt"
	"net/http"

	"github.com/mbier/aceleratomo/acelerato"
)

type HandlerQuadro struct {
}

func NewHandlerQuadro() *HandlerQuadro {
	return &HandlerQuadro{}
}

func (c *HandlerQuadro) QuadroTrack(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTrack())
}

func (c *HandlerQuadro) QuadroAdm(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroAdm())
}

func (c *HandlerQuadro) QuadroTMSWEB(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTMSWEB())
}

func (c *HandlerQuadro) QuadroSMONET(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMONET())
}

func (c *HandlerQuadro) QuadroSMOWEB(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMOWEB())
}

func (c *HandlerQuadro) QuadroSMOCTE(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMOCTE())
}

func (c *HandlerQuadro) QuadroDelhpi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroDelphi())
}

func QuadroGeral(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroGeral())
}
