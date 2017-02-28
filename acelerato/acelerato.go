package acelerato

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"strconv"
	"bytes"
)

func getDemandasTrack()[]Demanda {
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

func getDemandasTMSWEB()[]Demanda {
	client := &http.Client{}

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=15"
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

func getDemandasSMOWEB()[]Demanda {
	client := &http.Client{}

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=16"
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

func getDemandasSMOFrete()[]Demanda {
	client := &http.Client{}

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=22"
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

func getDemandasSMONET()[]Demanda {
	client := &http.Client{}

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=22"
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

func GerarQuadroTrack() string {
	demandas := getDemandasTrack()

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

func GerarQuadroTMSWEB() string {
	demandas := getDemandasTMSWEB()

	qtdBacklog := 0
	qtdBacklogProblema := 0
	qtdBacklogMelhoria := 0
	qtdTeste := 0
	qtdTesteProblema := 0
	qtdTesteMelhoria := 0

	for _, demanda := range demandas {
		if demanda.KanbanStatus.KanbanStatusKey == 19 || demanda.KanbanStatus.KanbanStatusKey == 20 {
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

func GerarQuadroSMONET() string {
	demandas := getDemandasSMONET()

	qtdBacklog := 0
	qtdBacklogProblema := 0
	qtdBacklogMelhoria := 0
	qtdTeste := 0
	qtdTesteProblema := 0
	qtdTesteMelhoria := 0

	for _, demanda := range demandas {
		if demanda.KanbanStatus.KanbanStatusKey == 19 || demanda.KanbanStatus.KanbanStatusKey == 20 {
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

func GerarQuadroSMOWEB() string {
	demandas := getDemandasSMOWEB()

	qtdBacklog := 0
	qtdBacklogProblema := 0
	qtdBacklogMelhoria := 0
	qtdTeste := 0
	qtdTesteProblema := 0
	qtdTesteMelhoria := 0

	for _, demanda := range demandas {
		if demanda.KanbanStatus.KanbanStatusKey == 19 || demanda.KanbanStatus.KanbanStatusKey == 20 {
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

func GerarQuadroGeral() string {
	demandasTrack := getDemandasTrack()

	qtdBacklog := 0
	qtdBacklogProblema := 0
	qtdBacklogMelhoria := 0
	qtdTeste := 0
	qtdTesteProblema := 0
	qtdTesteMelhoria := 0

	for _, demanda := range demandasTrack {
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

	buffer.WriteString("TRACK " + strconv.Itoa(qtdBacklogMelhoria) +" ")
	buffer.WriteString(strconv.Itoa(qtdBacklogProblema) +" ")
	buffer.WriteString(strconv.Itoa(qtdTeste) +" ")
	buffer.WriteString(strconv.Itoa(qtdBacklog) +" ")
	buffer.WriteString(strconv.FormatFloat(((float64(qtdBacklogMelhoria)/ float64(qtdBacklog)) * 100.0),'f', 2, 64 ) +" ")
	buffer.WriteString(strconv.FormatFloat(((float64(qtdBacklogProblema)/ float64(qtdBacklog)) * 100.0),'f', 2, 64 ) +" ")

	return buffer.String()
}
