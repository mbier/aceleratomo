package acelerato

// CategoriaDado dados de cada projeto
type CategoriaDado struct {
	QtdProblema int `json:"qtdProblema"`
	QtdBlocante int `json:"qtdBlocante"`
	QtdRecusado int `json:"qtdRecusado"`
	QtdMelhoria int `json:"qtdMelhoria"`
}

// NewCategoriaDado gerar um novo Categoria Dado
func NewCategoriaDado() CategoriaDado {
	qi := CategoriaDado{}
	qi.QtdProblema = 0
	qi.QtdBlocante = 0
	qi.QtdRecusado = 0
	qi.QtdMelhoria = 0

	return qi
}

// Total de damanda do projeto
func (qi CategoriaDado) Total() int {
	return qi.QtdMelhoria + qi.QtdProblema
}

// Merge demandas esperando merge
func (qi *CategoriaDado) Merge(itemMerge CategoriaDado) {
	qi.QtdProblema = itemMerge.QtdProblema
	qi.QtdBlocante = itemMerge.QtdBlocante
	qi.QtdRecusado = itemMerge.QtdRecusado
	qi.QtdMelhoria = itemMerge.QtdMelhoria
}
