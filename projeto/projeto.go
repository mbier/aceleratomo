package projeto

// Projeto do acelerato
type Projeto struct {
	Nome                    string
	URLAcelerato            string
	EmDesenvolvimentoFiltro []int
	AgMergeFiltro           []int
	AgTesteFiltro           []int
	EmTesteFiltro           []int
	Grupo                   Grupo
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