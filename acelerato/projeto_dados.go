package acelerato

import (
	"github.com/mbier/aceleratomo/projeto"
)

// ProjetoDado é objeto de retorno da demanda do acelerato
type ProjetoDado struct {
	Projeto           projeto.Projeto `json:"projeto"`
	AgAprovacao       CategoriaDado   `json:"agAprovacao"`
	Aprovado          CategoriaDado   `json:"aprovado"`
	EmDesenvolvimento CategoriaDado   `json:"emDesenvolvimento"`
	AgMerge           CategoriaDado   `json:"agMerge"`
	QtdImpedimento    int             `json:"qtdImpedimento"`
	AgTeste           CategoriaDado   `json:"agTeste"`
	EmTeste           CategoriaDado   `json:"emTeste"`
	Recusado          CategoriaDado   `json:"recusado"`
}

// ProjetosDado lista de quadro
type ProjetosDado []ProjetoDado

func (q ProjetosDado) Len() int {
	return len(q)
}

func (q ProjetosDado) Less(i, j int) bool {
	return q[i].Projeto.Nome < q[j].Projeto.Nome
}

func (q ProjetosDado) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// NewProjetoDado gerar um novo Quadro
func NewProjetoDado() ProjetoDado {
	q := ProjetoDado{}
	q.Aprovado = NewCategoriaDado()
	q.EmDesenvolvimento = NewCategoriaDado()
	q.AgMerge = NewCategoriaDado()
	q.QtdImpedimento = 0
	q.AgTeste = NewCategoriaDado()
	q.EmTeste = NewCategoriaDado()

	return q
}

// TotalAprovado retorna a soma de aprovados
func (q *ProjetoDado) TotalAprovado() int {
	return q.Aprovado.Total()
}

// TotalEmDesenvolvimento retorna a soma de Em Desenvolvimento
func (q *ProjetoDado) TotalEmDesenvolvimento() int {
	return q.EmDesenvolvimento.Total()
}

// TotalAgTeste retorna a soma de Ag Teste
func (q *ProjetoDado) TotalAgTeste() int {
	return q.AgTeste.Total()
}

// TotalEmTeste retorna a soma de Em Teste
func (q *ProjetoDado) TotalEmTeste() int {
	return q.EmTeste.Total()
}

// TotalBacklog retorna a soma de backlog
func (q *ProjetoDado) TotalBacklog() int {
	return q.TotalAprovado() + q.TotalEmDesenvolvimento()
}

// TotalBacklogP retorna a soma de backlog problema
func (q *ProjetoDado) TotalBacklogP() int {
	return q.Aprovado.QtdProblema + q.EmDesenvolvimento.QtdProblema
}

// TotalBacklogM retorna a soma de backlog melhoria
func (q *ProjetoDado) TotalBacklogM() int {
	return q.Aprovado.QtdMelhoria + q.EmDesenvolvimento.QtdMelhoria
}

// TotalTeste retorna a soma de teste
func (q *ProjetoDado) TotalTeste() int {
	return q.TotalAgTeste() + q.TotalEmTeste()
}

// TotalTesteP retorna a soma de teste problema
func (q *ProjetoDado) TotalTesteP() int {
	return q.AgTeste.QtdProblema + q.EmTeste.QtdProblema
}

// TotalTesteM retorna a soma de teste melhoria
func (q *ProjetoDado) TotalTesteM() int {
	return q.AgTeste.QtdMelhoria + q.EmTeste.QtdMelhoria
}

// TotalAgMerge retorna a soma de ag merge
func (q *ProjetoDado) TotalAgMerge() int {
	return q.AgMerge.Total()
}

// Total retorna a soma de todos os indicadores
func (q *ProjetoDado) Total() int {
	return q.TotalBacklog() + q.TotalAgMerge() + q.TotalTeste()
}
