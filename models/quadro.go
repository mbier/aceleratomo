package models

// Quadro é objeto de retorno da demanda do acelerato
type Quadro struct {
	NomeProjeto        string
	QtdBacklogProblema int
	QtdBacklogMelhoria int
	QtdTesteProblema   int
	QtdTesteMelhoria   int
	QtdAgMerge         int
	QtdImpedimento     int
}

// Quadros lista de quadro
type Quadros []Quadro

func (q Quadros) Len() int {
	return len(q)
}

func (q Quadros) Less(i, j int) bool {
	return q[i].NomeProjeto < q[j].NomeProjeto
}

func (q Quadros) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// NewQuadro gerar um novo Quadro
func NewQuadro() *Quadro {
	q := &Quadro{}
	q.QtdBacklogProblema = 0
	q.QtdBacklogMelhoria = 0
	q.QtdTesteProblema = 0
	q.QtdTesteMelhoria = 0
	q.QtdAgMerge = 0
	q.QtdImpedimento = 0
	return q
}

// Total retorna a soma de todos os indicadores
func (q *Quadro) Total() int {
	return q.QtdBacklogProblema + q.QtdBacklogMelhoria + q.QtdTesteProblema + q.QtdTesteMelhoria + q.QtdAgMerge + q.QtdImpedimento
}
