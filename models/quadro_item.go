package models

type QuadroItem struct {
	QtdProblema int
	QtdBlocante int
	QtdRecusado int
	QtdMelhoria int
}

// NewQuadroItem gerar um novo Quadro Item
func NewQuadroItem() QuadroItem {
	qi := QuadroItem{}
	qi.QtdProblema = 0
	qi.QtdBlocante = 0
	qi.QtdRecusado = 0
	qi.QtdMelhoria = 0

	return qi
}

func (qi QuadroItem) Total() int {
	return qi.QtdMelhoria + qi.QtdProblema
}

func (qi *QuadroItem) Merge(quadroItemMerge QuadroItem) {
	qi.QtdProblema = quadroItemMerge.QtdProblema
	qi.QtdBlocante = quadroItemMerge.QtdBlocante
	qi.QtdRecusado = quadroItemMerge.QtdRecusado
	qi.QtdMelhoria = quadroItemMerge.QtdMelhoria
}
