package models

type TipoProjeto int;

const (
	TRACK TipoProjeto = iota
	ADM
	SMO_CTE
	SMO_NET
	SMO_WEB
	TMS_WEB
)
