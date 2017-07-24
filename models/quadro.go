package models

import (
	"github.com/mbier/aceleratomo/projeto"
)

// Quadro é objeto de retorno da demanda do acelerato
type Quadro struct {
	Projeto           projeto.Projeto
	Aprovado          QuadroItem
	EmDesenvolvimento QuadroItem
	AgMerge           QuadroItem
	QtdImpedimento    int
	AgTeste           QuadroItem
	EmTeste           QuadroItem
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
func NewQuadro() Quadro {
	q := Quadro{}
	q.Aprovado = NewQuadroItem()
	q.EmDesenvolvimento = NewQuadroItem()
	q.AgMerge = NewQuadroItem()
	q.QtdImpedimento = 0
	q.AgTeste = NewQuadroItem()
	q.EmTeste = NewQuadroItem()

	return q
}

// TotalAprovado retorna a soma de aprovados
func (q *Quadro) TotalAprovado() int {
	return q.Aprovado.Total()
}

// TotalEmDesenvolvimento retorna a soma de Em Desenvolvimento
func (q *Quadro) TotalEmDesenvolvimento() int {
	return q.EmDesenvolvimento.Total()
}

// TotalAgTeste retorna a soma de Ag Teste
func (q *Quadro) TotalAgTeste() int {
	return q.AgTeste.Total()
}

// TotalEmTeste retorna a soma de Em Teste
func (q *Quadro) TotalEmTeste() int {
	return q.EmTeste.Total()
}

// TotalBacklog retorna a soma de backlog
func (q *Quadro) TotalBacklog() int {
	return q.TotalAprovado() + q.TotalEmDesenvolvimento()
}

// TotalBacklogP retorna a soma de backlog problema
func (q *Quadro) TotalBacklogP() int {
	return q.Aprovado.QtdProblema + q.EmDesenvolvimento.QtdProblema
}

// TotalBacklogM retorna a soma de backlog melhoria
func (q *Quadro) TotalBacklogM() int {
	return q.Aprovado.QtdMelhoria + q.EmDesenvolvimento.QtdMelhoria
}

// TotalTeste retorna a soma de teste
func (q *Quadro) TotalTeste() int {
	return q.TotalAgTeste() + q.TotalEmTeste()
}

// TotalTesteP retorna a soma de teste problema
func (q *Quadro) TotalTesteP() int {
	return q.AgTeste.QtdProblema + q.EmTeste.QtdProblema
}

// TotalTesteM retorna a soma de teste melhoria
func (q *Quadro) TotalTesteM() int {
	return q.AgTeste.QtdMelhoria + q.EmTeste.QtdMelhoria
}

// TotalAgMerge retorna a soma de ag merge
func (q *Quadro) TotalAgMerge() int {
	return q.AgMerge.Total()
}

// Total retorna a soma de todos os indicadores
func (q *Quadro) Total() int {
	return q.TotalBacklog() + q.TotalAgMerge() + q.TotalTeste()
}
