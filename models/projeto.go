package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Projeto struct {
		ID                 bson.ObjectId `bson:"_id"`
		Data               time.Time `bson:"data_geracao"`
		TipoProjeto        TipoProjeto `bson:"tipo_projeto"`
		QtdBacklogProblema int `bson:"qtd_backlog_problema"`
		QtdBacklogMelhoria int `bson:"qtd_backlog_melhoria"`
		QtdTesteProblema   int `bson:"qtd_teste_problema"`
		QtdTesteMelhoria   int `bson:"qtd_teste_melhoria"`
	}
)
