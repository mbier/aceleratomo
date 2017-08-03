package projeto

import (
	"errors"
	"sort"
	"strings"
)

// GetProjetos obtem os projetos do acelerato
func GetProjetos() Projetos {
	projetos := Projetos{}

	projetos = append(projetos, getProjetoADM())
	projetos = append(projetos, getProjetoDelphi())
	projetos = append(projetos, getProjetoLogrev())
	projetos = append(projetos, getProjetoPortalLogin())
	projetos = append(projetos, getProjetoSMOCTE())
	projetos = append(projetos, getProjetoSMOFRETE())
	projetos = append(projetos, getProjetoSMONET())
	projetos = append(projetos, getProjetoSMOTMS())
	projetos = append(projetos, getProjetoTMSWEB())

	sort.Sort(projetos)

	return projetos
}

// GetProjeto obtem o projeto a partir do nome
func GetProjeto(nome string) (*Projeto, error) {

	achou := false
	var projeto Projeto
	var err error

	for _, p := range GetProjetos() {
		if strings.ToLower(p.Nome) == strings.ToLower(nome) {
			projeto = p
			achou = true
		}
	}

	if !achou {
		err = errors.New("Projeto " + nome + " nao encontrado")
	}

	return &projeto, err
}

func getProjetoSMOFRETE() Projeto {
	projeto := Projeto{}
	projeto.Nome = "SMOFRETE"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=20,22"
	projeto.EmDesenvolvimentoFiltro = []int{9}
	projeto.AgMergeFiltro = []int{13}
	projeto.AgTesteFiltro = []int{10}
	projeto.EmTesteFiltro = []int{11}
	projeto.Grupo = FRETE

	return projeto
}

func getProjetoADM() Projeto {
	projeto := Projeto{}
	projeto.Nome = "ADM"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=26"
	projeto.EmDesenvolvimentoFiltro = []int{9}
	projeto.AgMergeFiltro = []int{13}
	projeto.AgTesteFiltro = []int{10}
	projeto.EmTesteFiltro = []int{11}
	projeto.Grupo = ADM

	return projeto
}

func getProjetoSMOTMS() Projeto {
	projeto := Projeto{}
	projeto.Nome = "SMOTMS"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=15"
	projeto.EmDesenvolvimentoFiltro = []int{18}
	projeto.AgTesteFiltro = []int{19}
	projeto.EmTesteFiltro = []int{20}
	projeto.AgMergeFiltro = []int{74}
	projeto.Grupo = FLEX

	return projeto
}

func getProjetoTMSWEB() Projeto {
	projeto := Projeto{}
	projeto.Nome = "TMSWEB"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=15"
	projeto.EmDesenvolvimentoFiltro = []int{18}
	projeto.AgMergeFiltro = []int{13}
	projeto.AgTesteFiltro = []int{10}
	projeto.EmTesteFiltro = []int{11}
	projeto.Grupo = TMS

	return projeto
}

func getProjetoSMONET() Projeto {
	projeto := Projeto{}
	projeto.Nome = "SMONET"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=22"
	projeto.EmDesenvolvimentoFiltro = []int{18}
	projeto.AgTesteFiltro = []int{19}
	projeto.EmTesteFiltro = []int{20}
	projeto.AgMergeFiltro = []int{74}
	projeto.Grupo = PORTAL

	return projeto
}

func getProjetoSMOCTE() Projeto {
	projeto := Projeto{}
	projeto.Nome = "SMOCTE"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=16"
	projeto.EmDesenvolvimentoFiltro = []int{18}
	projeto.AgTesteFiltro = []int{19}
	projeto.EmTesteFiltro = []int{20}
	projeto.AgMergeFiltro = []int{74}
	projeto.Grupo = FLEX

	return projeto
}

func getProjetoLogrev() Projeto {
	projeto := Projeto{}
	projeto.Nome = "Logrev"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=25"
	projeto.EmDesenvolvimentoFiltro = []int{9}
	projeto.AgMergeFiltro = []int{13}
	projeto.AgTesteFiltro = []int{10}
	projeto.EmTesteFiltro = []int{11}
	projeto.Grupo = PORTAL

	return projeto
}

func getProjetoDelphi() Projeto {
	projeto := Projeto{}
	projeto.Nome = "Delphi"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=5&categorias=17"
	projeto.EmDesenvolvimentoFiltro = []int{27}
	projeto.AgTesteFiltro = []int{28}
	projeto.EmTesteFiltro = []int{29}
	projeto.Grupo = DELPHI

	return projeto
}

func getProjetoPortalLogin() Projeto {
	projeto := Projeto{}
	projeto.Nome = "Portal"
	projeto.URLAcelerato = "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=18"
	projeto.EmDesenvolvimentoFiltro = []int{9}
	projeto.AgMergeFiltro = []int{13}
	projeto.AgTesteFiltro = []int{10}
	projeto.EmTesteFiltro = []int{11}
	projeto.Grupo = PORTAL

	return projeto
}
