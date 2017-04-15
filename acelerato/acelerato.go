package acelerato

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/mbier/aceleratomo/models"
)

// Acesso ao Acelerato
const (
	Usuario string = "marlon.bier@mosistemas.com"
	Token   string = "5cGDISvo1gM2Gi7tO7G+jA=="
)

func getDemandas(url string) []models.Demanda {
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

	var demandas []models.Demanda

	json.Unmarshal(responseData, &demandas)

	return demandas
}

func getDemandasTrack() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=20,22"

	return getDemandas(url)
}

func getDemandasAdm() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=11&categorias=26"

	return getDemandas(url)
}

func getDemandasTMSWEB() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=15"

	return getDemandas(url)
}

func getDemandasSMOWEB() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=16"

	return getDemandas(url)
}

func getDemandasSMONET() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=22"

	return getDemandas(url)
}

func getDemandasSMOCTE() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=16"

	return getDemandas(url)
}

func getDemandasLogrev() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=2&categorias=25"

	return getDemandas(url)
}

func getDemandasDelphi() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=5&categorias=17"

	return getDemandas(url)
}

// GerarQuadroTrack gera as informacoes do track
func GerarQuadroTrack() string {

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := GerarDadosQuadroTrack()

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

// GerarDadosQuadroTrack gera as informacoes do track
func GerarDadosQuadroTrack() (qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria int) {
	demandas := getDemandasTrack()

	testeFiltro := []int{10, 11}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria = gerarQuadro(demandas, testeFiltro)

	return
}

// GerarQuadroAdm gera as informacoes do track
func GerarQuadroAdm() string {
	demandas := getDemandasAdm()

	testeFiltro := []int{10, 11}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

// GerarQuadroTMSWEB gera as informacoes do track
func GerarQuadroTMSWEB() string {
	demandas := getDemandasTMSWEB()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

// GerarQuadroSMONET gera as informacoes do track
func GerarQuadroSMONET() string {
	demandas := getDemandasSMONET()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

// GerarQuadroSMOWEB gera as informacoes do track
func GerarQuadroSMOWEB() string {
	demandas := getDemandasSMOWEB()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

// GerarQuadroSMOCTE gera as informacoes do track
func GerarQuadroSMOCTE() string {
	demandas := getDemandasSMOCTE()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

// GerarQuadroDelphi gera as informacoes do track
func GerarQuadroDelphi() string {
	demandas := getDemandasDelphi()

	testeFiltro := []int{19, 28}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

// GerarQuadroGeral gera as informacoes do track
func GerarQuadroGeral() string {
	demandasSmocte := getDemandasSMOCTE()
	demandasTmsweb := getDemandasTMSWEB()
	demandasSmonet := getDemandasSMONET()
	demandasTrack := getDemandasTrack()
	demandasAdm := getDemandasAdm()
	demandasLogrev := getDemandasLogrev()
	demandasDelphi := getDemandasDelphi()

	testeFiltroTrack := []int{10, 11}
	testeFiltroSMO := []int{19, 20}
	testeFiltroDelphi := []int{19, 28}

	qtdBacklogProblemaTrack, qtdBacklogMelhoriaTrack, qtdTesteProblemaTrack, qtdTesteMelhoriaTrack := gerarQuadro(demandasTrack, testeFiltroTrack)
	qtdBacklogProblemaAdm, qtdBacklogMelhoriaAdm, qtdTesteProblemaAdm, qtdTesteMelhoriaAdm := gerarQuadro(demandasAdm, testeFiltroTrack)
	qtdBacklogProblemaTmsweb, qtdBacklogMelhoriaTmsweb, qtdTesteProblemaTmsweb, qtdTesteMelhoriaTmsweb := gerarQuadro(demandasTmsweb, testeFiltroSMO)
	qtdBacklogProblemaSmonet, qtdBacklogMelhoriaSmonet, qtdTesteProblemaSmonet, qtdTesteMelhoriaSmonet := gerarQuadro(demandasSmonet, testeFiltroSMO)
	qtdBacklogProblemaSmocte, qtdBacklogMelhoriaSmocte, qtdTesteProblemaSmocte, qtdTesteMelhoriaSmocte := gerarQuadro(demandasSmocte, testeFiltroSMO)
	qtdBacklogProblemaLogrev, qtdBacklogMelhoriaLogrev, qtdTesteProblemaLogrev, qtdTesteMelhoriaLogrev := gerarQuadro(demandasLogrev, testeFiltroTrack)
	qtdBacklogProblemaDelphi, qtdBacklogMelhoriaDelphi, qtdTesteProblemaDelphi, qtdTesteMelhoriaDelphi := gerarQuadro(demandasDelphi, testeFiltroDelphi)

	qtdBacklogProblemaGeral := qtdBacklogProblemaTrack + qtdBacklogProblemaTmsweb + qtdBacklogProblemaSmonet + qtdBacklogProblemaAdm + qtdBacklogProblemaSmocte + qtdBacklogProblemaLogrev + qtdBacklogProblemaDelphi
	qtdBacklogMelhoriaGeral := qtdBacklogMelhoriaTrack + qtdBacklogMelhoriaTmsweb + qtdBacklogMelhoriaSmonet + qtdBacklogMelhoriaAdm + qtdBacklogMelhoriaSmocte + qtdBacklogMelhoriaLogrev + qtdBacklogMelhoriaDelphi
	qtdTesteProblemaGeral := qtdTesteProblemaTrack + qtdTesteProblemaTmsweb + qtdTesteProblemaSmonet + qtdTesteProblemaAdm + qtdTesteProblemaSmocte + qtdTesteProblemaLogrev + qtdTesteProblemaDelphi
	qtdTesteMelhoriaGeral := qtdTesteMelhoriaTrack + qtdTesteMelhoriaTmsweb + qtdTesteMelhoriaSmonet + qtdTesteMelhoriaAdm + qtdTesteMelhoriaSmocte + qtdTesteMelhoriaLogrev + qtdTesteMelhoriaDelphi

	var buffer bytes.Buffer

	buffer.WriteString("<head>")
	buffer.WriteString("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css\" integrity=\"sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ\" crossorigin=\"anonymous\">")
	buffer.WriteString("<script src=\"https://code.jquery.com/jquery-3.1.1.slim.min.js\" integrity=\"sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js\" integrity=\"sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js\" integrity=\"sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("</head>")

	buffer.WriteString("<table style=\"width:100%\" class=\"table table-striped table-bordered\">")
	buffer.WriteString("<tr><th>Produto</th><th>Melhoria</th><th>Problema</th><th>AG. Teste</th><th>Total</th><th>&#37; Melhoria</th><th>&#37; Problema</th></tr>")

	buffer.WriteString(gerarQuadroGeralItem("SMOCTE", qtdBacklogProblemaSmocte, qtdBacklogMelhoriaSmocte, qtdTesteProblemaSmocte, qtdTesteMelhoriaSmocte))
	buffer.WriteString(gerarQuadroGeralItem("SMOTMS", qtdBacklogProblemaTmsweb, qtdBacklogMelhoriaTmsweb, qtdTesteProblemaTmsweb, qtdTesteMelhoriaTmsweb))
	buffer.WriteString(gerarQuadroGeralItem("ADM", qtdBacklogProblemaAdm, qtdBacklogMelhoriaAdm, qtdTesteProblemaAdm, qtdTesteMelhoriaAdm))
	buffer.WriteString(gerarQuadroGeralItem("LOGREV", qtdBacklogProblemaLogrev, qtdBacklogMelhoriaLogrev, qtdTesteProblemaLogrev, qtdTesteMelhoriaLogrev))
	buffer.WriteString(gerarQuadroGeralItem("SMONET", qtdBacklogProblemaSmonet, qtdBacklogMelhoriaSmonet, qtdTesteProblemaSmonet, qtdTesteMelhoriaSmonet))
	buffer.WriteString(gerarQuadroGeralItem("SMOFRETE", qtdBacklogProblemaTrack, qtdBacklogMelhoriaTrack, qtdTesteProblemaTrack, qtdTesteMelhoriaTrack))
	buffer.WriteString(gerarQuadroGeralItem("DELPHI", qtdBacklogProblemaDelphi, qtdBacklogMelhoriaDelphi, qtdTesteProblemaDelphi, qtdTesteMelhoriaDelphi))
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
	buffer.WriteString("<td>" + strconv.Itoa(qtdTesteProblema+qtdTesteMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdBacklogProblema+qtdBacklogMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.FormatFloat(((float64(qtdBacklogMelhoria)/float64(qtdBacklogProblema+qtdBacklogMelhoria))*100.0), 'f', 2, 64) + "</td>")
	buffer.WriteString("<td>" + strconv.FormatFloat(((float64(qtdBacklogProblema)/float64(qtdBacklogProblema+qtdBacklogMelhoria))*100.0), 'f', 2, 64) + "</td>")
	buffer.WriteString("</tr>")

	return buffer.String()
}

func gerarQuadro(demandas []models.Demanda, testeFilter []int) (qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria int) {
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

	buffer.WriteString("<head>")
	buffer.WriteString("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css\" integrity=\"sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ\" crossorigin=\"anonymous\">")
	buffer.WriteString("<script src=\"https://code.jquery.com/jquery-3.1.1.slim.min.js\" integrity=\"sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js\" integrity=\"sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js\" integrity=\"sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("</head>")

	buffer.WriteString("<table style=\"width:100%\" class=\"table table-striped table-bordered\">")
	buffer.WriteString("<tr><th></th><th>Total</th><th>Problema</th><th>Melhoria</th></tr>")

	buffer.WriteString("<tr><td>Backlog</td><td>" + strconv.Itoa(qtdBacklogProblema+qtdBacklogMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdBacklogProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdBacklogMelhoria) + "</td></tr>")

	buffer.WriteString("<tr><td>Em Teste</td><td>" + strconv.Itoa(qtdTesteProblema+qtdTesteMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdTesteProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdTesteMelhoria) + "</td></tr>")

	buffer.WriteString("<tr><td>Total</td><td>" + strconv.Itoa(qtdBacklogProblema+qtdBacklogMelhoria+qtdTesteProblema+qtdTesteMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdTesteProblema+qtdBacklogProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(qtdTesteMelhoria+qtdBacklogMelhoria) + "</td></tr>")

	buffer.WriteString("</table>")

	return buffer.String()
}
