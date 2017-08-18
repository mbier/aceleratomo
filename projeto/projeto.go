package projeto

// Projeto do acelerato
type Projeto struct {
	Nome                    string `json:"nome"`
	URLAcelerato            string `json:"urlAcelerato"`
	AprovadoFiltro          []int  `json:"aprocadoFiltro"`
	EmDesenvolvimentoFiltro []int  `json:"emDesenvolvimentoFiltro"`
	AgMergeFiltro           []int  `json:"agMergeFiltro"`
	AgTesteFiltro           []int  `json:"agTesteFiltro"`
	EmTesteFiltro           []int  `json:"emTesteFiltro"`
	RecusadoFiltro          []int  `json:"recusadoFiltro"`
	Grupo                   Grupo  `json:"grupo"`
}

// Projetos lista de projetos
type Projetos []Projeto

func (projetos Projetos) Len() int {
	return len(projetos)
}

func (projetos Projetos) Less(i, j int) bool {
	return projetos[i].Nome < projetos[j].Nome
}

func (projetos Projetos) Swap(i, j int) {
	projetos[i], projetos[j] = projetos[j], projetos[i]
}
