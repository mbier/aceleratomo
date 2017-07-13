package acelerato

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/mbier/aceleratomo/models"
	"github.com/mbier/aceleratomo/projeto"
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

func getDemandasProjeto(projeto *projeto.Projeto) []models.Demanda {

	return getDemandas(projeto.URLAcelerato)
}

// GerarQuadro com as informacoes do projeto
func GerarQuadro(projeto *projeto.Projeto) string {

	quadro := GerarDadosQuadro(projeto)

	return gerarQuadroString(quadro)
}

// GerarDadosQuadro gera as informacoes do track
func GerarDadosQuadro(projeto *projeto.Projeto) (quadro models.Quadro) {
	demandas := getDemandasProjeto(projeto)

	quadro = *gerarQuadro(demandas, projeto.TesteFiltro, projeto.AgMergeFiltro)

	return
}

// GerarQuadroGeral gera as informacoes do track
func GerarQuadroGeral() string {

	var qtdBacklogProblemaGeral, qtdBacklogMelhoriaGeral, qtdTesteProblemaGeral, qtdTesteMelhoriaGeral, qtdAgMergeGeral, qtdImpedimentoGeral int

	var buffer bytes.Buffer

	buffer.WriteString("<head>")
	buffer.WriteString("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css\" integrity=\"sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ\" crossorigin=\"anonymous\">")
	buffer.WriteString("<script src=\"https://code.jquery.com/jquery-3.1.1.slim.min.js\" integrity=\"sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js\" integrity=\"sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js\" integrity=\"sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("</head>")

	buffer.WriteString("<table style=\"width:100%\" class=\"table table-striped table-bordered\">")
	buffer.WriteString("<tr><th>Produto</th><th>Melhoria</th><th>Problema</th><th>AG. Merge</th><th>AG. Teste</th><th>Total</th><th>&#37; Melhoria</th><th>&#37; Problema</th><th style=\"width: 50px;\">Impedimento</th></tr>")

	projetos := projeto.GetProjetos()

	out := make(chan models.Quadro, len(projetos))
	wg := sync.WaitGroup{}

	for _, p := range projetos {
		wg.Add(1)
		go func(projeto projeto.Projeto) {
			quadro := GerarDadosQuadro(&projeto)
			quadro.NomeProjeto = projeto.Nome
			wg.Done()
			out <- quadro
		}(p)
	}

	wg.Wait()
	close(out)

	var quadros models.Quadros

	for n := range out {
		quadros = append(quadros, n)
	}

	sort.Sort(quadros)

	for _, q := range quadros {

		buffer.WriteString(gerarQuadroGeralItem(q.NomeProjeto, &q))

		qtdBacklogProblemaGeral += q.QtdBacklogProblema
		qtdBacklogMelhoriaGeral += q.QtdBacklogMelhoria
		qtdTesteProblemaGeral += q.QtdTesteProblema
		qtdTesteMelhoriaGeral += q.QtdTesteMelhoria
		qtdAgMergeGeral += q.QtdAgMerge
		qtdImpedimentoGeral += q.QtdImpedimento
	}

	quadroGeral := models.NewQuadro()
	quadroGeral.QtdBacklogMelhoria = qtdBacklogMelhoriaGeral
	quadroGeral.QtdBacklogProblema = qtdBacklogProblemaGeral
	quadroGeral.QtdTesteMelhoria = qtdTesteProblemaGeral
	quadroGeral.QtdTesteProblema = qtdTesteMelhoriaGeral
	quadroGeral.QtdAgMerge = qtdAgMergeGeral
	quadroGeral.QtdImpedimento = qtdImpedimentoGeral

	buffer.WriteString(gerarQuadroGeralItem("Total", quadroGeral))

	buffer.WriteString("</table>")

	return buffer.String()
}

func gerarQuadroGeralItem(produto string, quadro *models.Quadro) string {
	var buffer bytes.Buffer

	buffer.WriteString("<tr>")
	buffer.WriteString("<td>" + produto + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdBacklogMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdBacklogProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdAgMerge) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdTesteProblema+quadro.QtdTesteMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdBacklogProblema+quadro.QtdBacklogMelhoria) + "</td>")
	if quadro.QtdBacklogProblema+quadro.QtdBacklogMelhoria > 0 {
		buffer.WriteString("<td>" + strconv.FormatFloat(((float64(quadro.QtdBacklogMelhoria)/float64(quadro.QtdBacklogProblema+quadro.QtdBacklogMelhoria))*100.0), 'f', 2, 64) + "</td>")
	} else {
		buffer.WriteString("<td>0.00</td>")
	}
	if quadro.QtdBacklogProblema+quadro.QtdBacklogMelhoria > 0 {
		buffer.WriteString("<td>" + strconv.FormatFloat(((float64(quadro.QtdBacklogProblema)/float64(quadro.QtdBacklogProblema+quadro.QtdBacklogMelhoria))*100.0), 'f', 2, 64) + "</td>")
	} else {
		buffer.WriteString("<td>0.00</td>")
	}
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdImpedimento) + "</td>")
	buffer.WriteString("</tr>")

	return buffer.String()
}

func gerarQuadro(demandas []models.Demanda, testeFilter, agMergeFilter []int) *models.Quadro {

	quadro := models.NewQuadro()

	for _, demanda := range demandas {
		if demanda.UsuarioImpedimento.Nome != "" {
			quadro.QtdImpedimento++
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, testeFilter) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.QtdTesteProblema++
			} else {
				quadro.QtdTesteMelhoria++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, agMergeFilter) {
			quadro.QtdAgMerge++
		} else {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.QtdBacklogProblema++
			} else {
				quadro.QtdBacklogMelhoria++
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

func gerarQuadroString(quadro models.Quadro) string {
	var buffer bytes.Buffer

	buffer.WriteString("<head>")
	buffer.WriteString("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css\" integrity=\"sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ\" crossorigin=\"anonymous\">")
	buffer.WriteString("<script src=\"https://code.jquery.com/jquery-3.1.1.slim.min.js\" integrity=\"sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js\" integrity=\"sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js\" integrity=\"sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("</head>")

	buffer.WriteString("<table style=\"width:100%\" class=\"table table-striped table-bordered\">")
	buffer.WriteString("<tr><th></th><th>Total</th><th>Problema</th><th>Melhoria</th></tr>")

	buffer.WriteString("<tr><td>Backlog</td><td>" + strconv.Itoa(quadro.QtdBacklogProblema+quadro.QtdBacklogMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdBacklogProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdBacklogMelhoria) + "</td></tr>")

	if quadro.QtdAgMerge > 0 {
		buffer.WriteString("<tr><td>Ag Merge</td><td>" + strconv.Itoa(quadro.QtdAgMerge) + "</td>")
		buffer.WriteString("<td></td>")
		buffer.WriteString("<td></td></tr>")
	}

	if quadro.QtdImpedimento > 0 {
		buffer.WriteString("<tr><td>Impedimento</td><td>" + strconv.Itoa(quadro.QtdImpedimento) + "</td>")
		buffer.WriteString("<td></td>")
		buffer.WriteString("<td></td></tr>")
	}

	buffer.WriteString("<tr><td>Em Teste</td><td>" + strconv.Itoa(quadro.QtdTesteProblema+quadro.QtdTesteMelhoria) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdTesteProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdTesteMelhoria) + "</td></tr>")

	buffer.WriteString("<tr><td>Total</td><td>" + strconv.Itoa(quadro.QtdBacklogProblema+quadro.QtdBacklogMelhoria+quadro.QtdTesteProblema+quadro.QtdTesteMelhoria+quadro.QtdAgMerge) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdTesteProblema+quadro.QtdBacklogProblema) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdTesteMelhoria+quadro.QtdBacklogMelhoria) + "</td></tr>")

	buffer.WriteString("</table>")

	return buffer.String()
}
