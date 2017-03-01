package acelerato

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"strconv"
	"bytes"
)

func getDemandasTrack() []Demanda {
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

func getDemandasAdm() []Demanda {
	client := &http.Client{}

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=11&categorias=26"
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

func getDemandasTMSWEB() []Demanda {
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

func getDemandasSMOWEB() []Demanda {
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

func getDemandasSMOFrete() []Demanda {
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

func getDemandasSMONET() []Demanda {
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

	testeFiltro := []int{10, 11}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroAdm() string {
	demandas := getDemandasAdm()

	testeFiltro := []int{10, 11}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroTMSWEB() string {
	demandas := getDemandasTMSWEB()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroSMONET() string {
	demandas := getDemandasSMONET()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroSMOWEB() string {
	demandas := getDemandasSMOWEB()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroGeral() string {
	demandasTmsweb := getDemandasTMSWEB()
	demandasSmonet := getDemandasSMONET()
	demandasTrack := getDemandasTrack()
	demandasAdm := getDemandasAdm()

	testeFiltroTrack := []int{10, 11}
	testeFiltroFlex := []int{10, 11}

	qtdBacklogProblemaTrack, qtdBacklogMelhoriaTrack, qtdTesteProblemaTrack, qtdTesteMelhoriaTrack := gerarQuadro(demandasTrack, testeFiltroTrack)
	qtdBacklogProblemaAdm, qtdBacklogMelhoriaAdm, qtdTesteProblemaAdm, qtdTesteMelhoriaAdm := gerarQuadro(demandasAdm, testeFiltroTrack)
	qtdBacklogProblemaTmsweb, qtdBacklogMelhoriaTmsweb, qtdTesteProblemaTmsweb, qtdTesteMelhoriaTmsweb := gerarQuadro(demandasTmsweb, testeFiltroFlex)
	qtdBacklogProblemaSmonet, qtdBacklogMelhoriaSmonet, qtdTesteProblemaSmonet, qtdTesteMelhoriaSmonet := gerarQuadro(demandasSmonet, testeFiltroFlex)

	qtdBacklogProblema := 0
	qtdBacklogMelhoria := 0
	qtdTesteProblema := 0
	qtdTesteMelhoria := 0

	qtdBacklogProblemaGeral := qtdBacklogProblemaTrack + qtdBacklogProblemaTmsweb + qtdBacklogProblemaSmonet+qtdBacklogProblemaAdm
	qtdBacklogMelhoriaGeral := qtdBacklogMelhoriaTrack + qtdBacklogMelhoriaTmsweb + qtdBacklogMelhoriaSmonet+qtdBacklogMelhoriaAdm
	qtdTesteProblemaGeral := qtdTesteProblemaTrack + qtdTesteProblemaTmsweb + qtdTesteProblemaSmonet+qtdTesteProblemaAdm
	qtdTesteMelhoriaGeral := qtdTesteMelhoriaTrack + qtdTesteMelhoriaTmsweb + qtdTesteMelhoriaSmonet+qtdTesteMelhoriaAdm

	var buffer bytes.Buffer

	buffer.WriteString("<table style=\"width:100%\">")
	buffer.WriteString("<tr><th>Produto</th><th>Melhoria</th><th>Problema</th><th>AG. Teste</th><th>Total</th><th>&#37; Melhoria</th><th>&#37; Problema</th></tr>")

	buffer.WriteString(gerarQuadroGeralItem("SMOCTE", qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria))
	buffer.WriteString(gerarQuadroGeralItem("SMOTMS", qtdBacklogProblemaTmsweb, qtdBacklogMelhoriaTmsweb, qtdTesteProblemaTmsweb, qtdTesteMelhoriaTmsweb))
	buffer.WriteString(gerarQuadroGeralItem("ADM", qtdBacklogProblemaAdm, qtdBacklogMelhoriaAdm, qtdTesteProblemaAdm, qtdTesteMelhoriaAdm))
	buffer.WriteString(gerarQuadroGeralItem("SMO-FRETE", qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria))
	buffer.WriteString(gerarQuadroGeralItem("LOGREV", qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria))
	buffer.WriteString(gerarQuadroGeralItem("SMONET", qtdBacklogProblemaSmonet, qtdBacklogMelhoriaSmonet, qtdTesteProblemaSmonet, qtdTesteMelhoriaSmonet))
	buffer.WriteString(gerarQuadroGeralItem("TRACK", qtdBacklogProblemaTrack, qtdBacklogMelhoriaTrack, qtdTesteProblemaTrack, qtdTesteMelhoriaTrack))
	buffer.WriteString(gerarQuadroGeralItem("Total", qtdBacklogProblemaGeral, qtdBacklogMelhoriaGeral, qtdTesteProblemaGeral, qtdTesteMelhoriaGeral))

	buffer.WriteString("</table>")

	return buffer.String()
}

func gerarQuadroGeralItem(produto string, qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria int) string {
	var buffer bytes.Buffer

	buffer.WriteString("<tr>")
	buffer.WriteString("<td>" + produto + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdBacklogMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdBacklogProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdTesteProblema + qtdTesteMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdBacklogProblema + qtdBacklogMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.FormatFloat(((float64(qtdBacklogMelhoria) / float64(qtdBacklogProblema + qtdBacklogMelhoria)) * 100.0), 'f', 2, 64) + "</td>")
	buffer.WriteString("<td>" + strconv.FormatFloat(((float64(qtdBacklogProblema) / float64(qtdBacklogProblema + qtdBacklogMelhoria)) * 100.0), 'f', 2, 64) + "</td>")
	buffer.WriteString("</tr>")

	return buffer.String()
}

func gerarQuadro(demandas []Demanda, testeFilter[]int) (qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria int) {
	qtdBacklogProblema = 0
	qtdBacklogMelhoria = 0
	qtdTesteProblema = 0
	qtdTesteMelhoria = 0

	for _, demanda := range demandas {
		if arrayContains(demanda.KanbanStatus.KanbanStatusKey, testeFilter) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				qtdTesteProblema++
			} else {
				qtdTesteMelhoria++
			}
		} else {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				qtdBacklogProblema++
			} else {
				qtdBacklogMelhoria++
			}
		}
	}
	return
}

func arrayContains(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria int) string {
	var buffer bytes.Buffer

	buffer.WriteString("Total em Backlog..........:" + strconv.Itoa(qtdBacklogProblema + qtdBacklogMelhoria) + "\n")
	buffer.WriteString("Total em Backlog Problema.:" + strconv.Itoa(qtdBacklogProblema) + "\n")
	buffer.WriteString("Total em Backlog Melhoria.:" + strconv.Itoa(qtdBacklogMelhoria) + "\n")

	buffer.WriteString("Total em Teste............:" + strconv.Itoa(qtdTesteProblema + qtdTesteMelhoria) + "\n")
	buffer.WriteString("Total em Teste Problema...:" + strconv.Itoa(qtdTesteProblema) + "\n")
	buffer.WriteString("Total em Teste Melhoria...:" + strconv.Itoa(qtdTesteMelhoria) + "\n")

	buffer.WriteString("Total.....................:" + strconv.Itoa(qtdBacklogProblema + qtdBacklogMelhoria + qtdTesteProblema + qtdTesteMelhoria) + "\n")
	buffer.WriteString("Total Problema............:" + strconv.Itoa(qtdTesteProblema + qtdBacklogProblema) + "\n")
	buffer.WriteString("Total Melhoria............:" + strconv.Itoa(qtdTesteMelhoria + qtdBacklogMelhoria) + "\n")

	return buffer.String()
}