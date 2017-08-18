package acelerato

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/mbier/aceleratomo/projeto"
)

// Acesso ao Acelerato
const (
	Usuario string = "marlon.bier@mosistemas.com"
	Token   string = "5cGDISvo1gM2Gi7tO7G+jA=="
)

func getDemandas(url string) []Demanda {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(Usuario, Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()

	var demandas []Demanda

	json.Unmarshal(responseData, &demandas)

	return demandas
}

func getDemandasProjeto(projeto *projeto.Projeto) []Demanda {

	return getDemandas(projeto.URLAcelerato)
}

// GerarDadosQuadro gera as informacoes do projeto
func GerarDadosQuadro(projeto *projeto.Projeto) (quadro ProjetoDado) {
	demandas := getDemandasProjeto(projeto)

	quadro = gerarQuadro(demandas, *projeto)

	quadro.Projeto = *projeto

	return
}

func gerarQuadro(demandas []Demanda, projeto projeto.Projeto) ProjetoDado {

	quadro := NewProjetoDado()

	for _, demanda := range demandas {

		isBlocante := demanda.TipoDePrioridade.Descricao == "Blocante"

		isRecusado := false
		for _, tag := range demanda.Tags {
			if tag.Tag == "recusado" {
				isRecusado = true
				break
			}
		}

		if demanda.UsuarioImpedimento.Nome != "" {
			quadro.QtdImpedimento++
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.AgTesteFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.AgTeste.QtdProblema++
			} else {
				quadro.AgTeste.QtdMelhoria++
			}
			if isBlocante {
				quadro.AgTeste.QtdBlocante++
			}
			if isRecusado {
				quadro.AgTeste.QtdRecusado++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.EmTesteFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.EmTeste.QtdProblema++
			} else {
				quadro.EmTeste.QtdMelhoria++
			}
			if isBlocante {
				quadro.EmTeste.QtdBlocante++
			}
			if isRecusado {
				quadro.EmTeste.QtdRecusado++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.EmDesenvolvimentoFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.EmDesenvolvimento.QtdProblema++
			} else {
				quadro.EmDesenvolvimento.QtdMelhoria++
			}
			if isBlocante {
				quadro.EmDesenvolvimento.QtdBlocante++
			}
			if isRecusado {
				quadro.EmDesenvolvimento.QtdRecusado++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.AgMergeFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.AgMerge.QtdProblema++
			} else {
				quadro.AgMerge.QtdMelhoria++
			}
			if isBlocante {
				quadro.AgMerge.QtdBlocante++
			}
			if isRecusado {
				quadro.AgMerge.QtdRecusado++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.AprovadoFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.Aprovado.QtdProblema++
			} else {
				quadro.Aprovado.QtdMelhoria++
			}
			if isBlocante {
				quadro.Aprovado.QtdBlocante++
			}
			if isRecusado {
				quadro.Aprovado.QtdRecusado++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.RecusadoFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.Recusado.QtdProblema++
			} else {
				quadro.Recusado.QtdMelhoria++
			}
			if isBlocante {
				quadro.Recusado.QtdBlocante++
			}
			if isRecusado {
				quadro.Recusado.QtdRecusado++
			}
		} else {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.AgAprovacao.QtdProblema++
			} else {
				quadro.AgAprovacao.QtdMelhoria++
			}
			if isBlocante {
				quadro.AgAprovacao.QtdBlocante++
			}
			if isRecusado {
				quadro.AgAprovacao.QtdRecusado++
			}
		}
	}

	return quadro
}

// GerarDadosQuadros gera as informacoes dos projetos
func GerarDadosQuadros() ProjetosDado {
	projetos := projeto.GetProjetos()

	ch := make(chan ProjetoDado, len(projetos))
	wg := sync.WaitGroup{}

	for _, p := range projetos {
		wg.Add(1)
		go func(projeto projeto.Projeto) {
			quadro := GerarDadosQuadro(&projeto)
			quadro.Projeto = projeto

			defer wg.Done()

			ch <- quadro
		}(p)
	}

	wg.Wait()
	close(ch)
	var quadros ProjetosDado

	for n := range ch {
		quadros = append(quadros, n)
	}

	sort.Sort(quadros)

	return quadros
}

func arrayContains(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GerarDadosQuadrosTeste gera as informacoes dos projetos
func GerarDadosQuadrosTeste() ProjetosDado {

	quadros := GerarDadosQuadros()

	grupoProjetos := make(map[projeto.Grupo]ProjetoDado)

	for _, q := range quadros {

		quadroGrupo := grupoProjetos[q.Projeto.Grupo]

		quadroGrupo.Projeto.Nome = q.Projeto.Grupo.ToString()

		quadroGrupo.AgAprovacao.Merge(q.AgAprovacao)
		quadroGrupo.Aprovado.Merge(q.Aprovado)
		quadroGrupo.EmDesenvolvimento.Merge(q.EmDesenvolvimento)
		quadroGrupo.AgMerge.Merge(q.AgMerge)
		quadroGrupo.AgTeste.Merge(q.AgTeste)
		quadroGrupo.EmTeste.Merge(q.EmTeste)
		quadroGrupo.Recusado.Merge(q.Recusado)
		quadroGrupo.QtdImpedimento += q.QtdImpedimento

		grupoProjetos[q.Projeto.Grupo] = quadroGrupo
	}

	retorno := []ProjetoDado{}
	for gp := range grupoProjetos {
		dados := grupoProjetos[gp]
		dados.Projeto.Nome = gp.ToString()
		dados.Projeto.Grupo = gp

		retorno = append(retorno, dados)
	}

	return retorno
}
