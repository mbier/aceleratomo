package acelerato

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

// GerarDadosQuadro gera as informacoes do track
func GerarDadosQuadro(projeto *projeto.Projeto) (quadro ProjetoDado) {
	demandas := getDemandasProjeto(projeto)

	quadro = gerarQuadro(demandas, *projeto)

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
		} else {
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
		}
	}

	return quadro
}

func arrayContains(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
