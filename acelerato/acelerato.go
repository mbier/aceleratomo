package acelerato

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"strconv"
	"bytes"
)

func getDemandas()[]Demanda {
	client := &http.Client{}

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=20,22"
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("marlon.bier@mosistemas.com", "5cGDISvo1gM2Gi7tO7G+jA==")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var demandas []Demanda

	json.Unmarshal(responseData, &demandas)

	return demandas
}

func GerarQuadro() string {
	demandas := getDemandas()

	qtdBacklog := 0
	qtdBacklogProblema := 0
	qtdBacklogMelhoria := 0
	qtdTeste := 0
	qtdTesteProblema := 0
	qtdTesteMelhoria := 0

	for _, demanda := range demandas {
		if demanda.KanbanStatus.KanbanStatusKey == 10 || demanda.KanbanStatus.KanbanStatusKey == 11 {
			qtdTeste++
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				qtdTesteProblema++
			} else {
				qtdTesteMelhoria++
			}
		} else {
			qtdBacklog++
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				qtdBacklogProblema++
			} else {
				qtdBacklogMelhoria++
			}
		}
	}

	var buffer bytes.Buffer

	buffer.WriteString("Total em Backlog..........:" + strconv.Itoa(qtdBacklog) + "\n")
	buffer.WriteString("Total em Backlog Problema.:" + strconv.Itoa(qtdBacklogProblema) + "\n")
	buffer.WriteString("Total em Backlog Melhoria.:" + strconv.Itoa(qtdBacklogMelhoria) + "\n")

	buffer.WriteString("Total em Teste............:" + strconv.Itoa(qtdTeste) + "\n")
	buffer.WriteString("Total em Teste Problema...:" + strconv.Itoa(qtdTesteProblema) + "\n")
	buffer.WriteString("Total em Teste Melhoria...:" + strconv.Itoa(qtdTesteMelhoria) + "\n")

	buffer.WriteString("Total.....................:" + strconv.Itoa(qtdTeste+qtdBacklog) + "\n")
	buffer.WriteString("Total Problema............:" + strconv.Itoa(qtdTesteProblema+qtdBacklogProblema) + "\n")
	buffer.WriteString("Total Melhoria............:" + strconv.Itoa(qtdTesteMelhoria+qtdBacklogMelhoria) + "\n")

	return buffer.String()
}
