package handlers

import (
	"fmt"
	"github.com/mbier/aceleratomo/acelerato"
	"net/http"
	"gopkg.in/mgo.v2"
)

type HandlerQuadro struct {
	mongo *mgo.Session
}

func NewHandlerQuadro(mongo *mgo.Session) *HandlerQuadro {
	return &HandlerQuadro{mongo: mongo}
}

func (c *HandlerQuadro) QuadroTrack(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTrack(c.mongo))
}

func (c *HandlerQuadro) QuadroAdm(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroAdm(c.mongo))
}

func (c *HandlerQuadro) QuadroTMSWEB(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroTMSWEB(c.mongo))
}

func (c *HandlerQuadro) QuadroSMONET(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMONET(c.mongo))
}

func (c *HandlerQuadro) QuadroSMOWEB(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMOWEB(c.mongo))
}

func (c *HandlerQuadro) QuadroSMOCTE(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroSMOCTE(c.mongo))
}

func QuadroGeral(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, acelerato.GerarQuadroGeral())
}