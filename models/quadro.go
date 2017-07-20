package models

import (
	"github.com/mbier/aceleratomo/projeto"
)

// Quadro é objeto de retorno da demanda do acelerato
type Quadro struct {
	Projeto               projeto.Projeto
	QtdAprovadoP          int
	QtdAprovadoM          int
	QtdEmDesenvolvimentoP int
	QtdEmDesenvolvimentoM int
	QtdAgMerge            int
	QtdImpedimento        int
	QtdAgTesteP           int
	QtdAgTestePB          int
	QtdAgTesteM           int
	QtdEmTesteP           int
	QtdEmTestePB          int
	QtdEmTesteM           int
	QtdRecusado           int
	QtdBlocante           int
}

// Quadros lista de quadro
type Quadros []Quadro

func (q Quadros) Len() int {
	return len(q)
}

func (q Quadros) Less(i, j int) bool {
	return q[i].Projeto.Nome < q[j].Projeto.Nome
}

func (q Quadros) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// NewQuadro gerar um novo Quadro
func NewQuadro() *Quadro {
	q := &Quadro{}
	q.QtdAprovadoP = 0
	q.QtdAprovadoM = 0
	q.QtdEmDesenvolvimentoP = 0
	q.QtdEmDesenvolvimentoM = 0
	q.QtdAgMerge = 0
	q.QtdImpedimento = 0
	q.QtdAgTesteP = 0
	q.QtdAgTesteM = 0
	q.QtdEmTesteP = 0
	q.QtdEmTesteM = 0

	return q
}

// TotalAprovado retorna a soma de aprovados
func (q *Quadro) TotalAprovado() int {
	return q.QtdAprovadoP + q.QtdAprovadoM
}

// TotalEmDesenvolvimento retorna a soma de Em Desenvolvimento
func (q *Quadro) TotalEmDesenvolvimento() int {
	return q.QtdEmDesenvolvimentoP + q.QtdEmDesenvolvimentoM
}

// TotalAgTeste retorna a soma de Ag Teste
func (q *Quadro) TotalAgTeste() int {
	return q.QtdAgTesteP + q.QtdAgTestePB + q.QtdAgTesteM
}

// TotalEmTeste retorna a soma de Em Teste
func (q *Quadro) TotalEmTeste() int {
	return q.QtdEmTesteP + q.QtdEmTestePB + q.QtdEmTesteM
}

// TotalBacklog retorna a soma de backlog
func (q *Quadro) TotalBacklog() int {
	return q.TotalAprovado() + q.TotalEmDesenvolvimento()
}

// TotalBacklogP retorna a soma de backlog problema
func (q *Quadro) TotalBacklogP() int {
	return q.QtdAprovadoP + q.QtdEmDesenvolvimentoP
}

// TotalBacklogM retorna a soma de backlog melhoria
func (q *Quadro) TotalBacklogM() int {
	return q.QtdAprovadoM + q.QtdEmDesenvolvimentoM
}

// TotalTeste retorna a soma de teste
func (q *Quadro) TotalTeste() int {
	return q.TotalAgTeste() + q.TotalEmTeste()
}

// TotalTesteP retorna a soma de teste problema
func (q *Quadro) TotalTesteP() int {
	return q.QtdAgTesteP + q.QtdEmTesteP
}

// TotalTesteM retorna a soma de teste melhoria
func (q *Quadro) TotalTesteM() int {
	return q.QtdAgTesteM + q.QtdEmTesteM
}

// Total retorna a soma de todos os indicadores
func (q *Quadro) Total() int {
	return q.TotalBacklog() + q.QtdAgMerge + q.TotalTeste()
}
