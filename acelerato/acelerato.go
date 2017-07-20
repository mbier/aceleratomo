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

	quadro = *gerarQuadro(demandas, *projeto)

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
			quadro.Projeto = projeto
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

		buffer.WriteString(gerarQuadroGeralItem(q.Projeto.Nome, &q))

		qtdBacklogProblemaGeral += q.TotalBacklogP()
		qtdBacklogMelhoriaGeral += q.TotalBacklogM()
		qtdTesteProblemaGeral += q.TotalTesteP()
		qtdTesteMelhoriaGeral += q.TotalTesteM()
		qtdAgMergeGeral += q.QtdAgMerge
		qtdImpedimentoGeral += q.QtdImpedimento
	}

	quadroGeral := models.NewQuadro()
	quadroGeral.QtdAprovadoM = qtdBacklogMelhoriaGeral
	quadroGeral.QtdAprovadoP = qtdBacklogProblemaGeral
	quadroGeral.QtdEmTesteP = qtdTesteProblemaGeral
	quadroGeral.QtdEmTesteM = qtdTesteMelhoriaGeral
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
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogM()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogP()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdAgMerge) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalTeste()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklog()) + "</td>")
	if quadro.TotalBacklog() > 0 {
		buffer.WriteString("<td>" + strconv.FormatFloat(((float64(quadro.TotalBacklogM())/float64(quadro.TotalBacklog()))*100.0), 'f', 2, 64) + "</td>")
		buffer.WriteString("<td>" + strconv.FormatFloat(((float64(quadro.TotalBacklogP())/float64(quadro.TotalBacklog()))*100.0), 'f', 2, 64) + "</td>")
	} else {
		buffer.WriteString("<td>0.00</td>")
		buffer.WriteString("<td>0.00</td>")
	}
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdImpedimento) + "</td>")
	buffer.WriteString("</tr>")

	return buffer.String()
}

func gerarQuadro(demandas []models.Demanda, projeto projeto.Projeto) *models.Quadro {

	quadro := models.NewQuadro()

	for _, demanda := range demandas {
		if demanda.UsuarioImpedimento.Nome != "" {
			quadro.QtdImpedimento++
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.AgTesteFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				if demanda.TipoDePrioridade.Descricao == "Blocante" {
					quadro.QtdAgTestePB++
				} else {
					quadro.QtdAgTesteP++
				}
			} else {
				quadro.QtdAgTesteM++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.EmTesteFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				if demanda.TipoDePrioridade.Descricao == "Blocante" {
					quadro.QtdEmTestePB++
				} else {
					quadro.QtdEmTesteP++
				}
			} else {
				quadro.QtdEmTesteM++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.EmDesenvolvimentoFiltro) {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.QtdEmDesenvolvimentoP++
			} else {
				quadro.QtdEmDesenvolvimentoM++
			}
		} else if arrayContains(demanda.KanbanStatus.KanbanStatusKey, projeto.AgMergeFiltro) {
			quadro.QtdAgMerge++
		} else {
			if demanda.TipoDeTicket.TipoDeTicketKey == 3 {
				quadro.QtdAprovadoP++
			} else {
				quadro.QtdAprovadoM++
			}
		}

		for _, tag := range demanda.Tags {
			if tag.Tag == "recusado" {
				quadro.QtdRecusado++
				break
			}
		}

		if demanda.TipoDePrioridade.Descricao == "Blocante" {
			quadro.QtdBlocante++
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

	buffer.WriteString("<tr><td>Backlog</td><td>" + strconv.Itoa(quadro.TotalBacklog()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogP()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogM()) + "</td></tr>")

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

	buffer.WriteString("<tr><td>Em Teste</td><td>" + strconv.Itoa(quadro.TotalTeste()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalTesteP()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalTesteM()) + "</td></tr>")

	buffer.WriteString("<tr><td>Total</td><td>" + strconv.Itoa(quadro.Total()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogP()+quadro.TotalTesteP()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogM()+quadro.TotalTesteM()) + "</td></tr>")

	buffer.WriteString("</table>")

	return buffer.String()
}

// GerarQuadroTestes gera as informacoes do track
func GerarQuadroTestes() string {

	var buffer bytes.Buffer

	buffer.WriteString("<head>")
	buffer.WriteString("<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css\" integrity=\"sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ\" crossorigin=\"anonymous\">")
	buffer.WriteString("<script src=\"https://code.jquery.com/jquery-3.1.1.slim.min.js\" integrity=\"sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js\" integrity=\"sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js\" integrity=\"sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn\" crossorigin=\"anonymous\"></script>")
	buffer.WriteString("</head>")

	buffer.WriteString("<table style=\"width:100%\" class=\"table table-striped table-bordered\">")
	buffer.WriteString("<tr><th>Equipe</th><th>Desenv</th><th>AG. Merge</th><th>AG. Teste</th><th>Em Teste</th><th>Recusado</th><th>Blocante</th></tr>")

	projetos := projeto.GetProjetos()

	out := make(chan models.Quadro, len(projetos))
	wg := sync.WaitGroup{}

	for _, p := range projetos {
		wg.Add(1)
		go func(projeto projeto.Projeto) {
			quadro := GerarDadosQuadro(&projeto)
			quadro.Projeto = projeto
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

	grupoProjetos := make(map[projeto.Grupo]models.Quadro)

	for _, q := range quadros {

		quadroGrupo := grupoProjetos[q.Projeto.Grupo]

		quadroGrupo.Projeto.Nome = q.Projeto.Grupo.ToString()

		quadroGrupo.QtdAprovadoP += q.QtdAprovadoP
		quadroGrupo.QtdAprovadoM += q.QtdAprovadoM
		quadroGrupo.QtdEmDesenvolvimentoP += q.QtdEmDesenvolvimentoP
		quadroGrupo.QtdEmDesenvolvimentoM += q.QtdEmDesenvolvimentoM
		quadroGrupo.QtdAgMerge += q.QtdAgMerge
		quadroGrupo.QtdImpedimento += q.QtdImpedimento
		quadroGrupo.QtdAgTesteP += q.QtdAgTesteP
		quadroGrupo.QtdAgTestePB += q.QtdAgTestePB
		quadroGrupo.QtdAgTesteM += q.QtdAgTesteM
		quadroGrupo.QtdEmTesteP += q.QtdEmTesteP
		quadroGrupo.QtdEmTestePB += q.QtdEmTestePB
		quadroGrupo.QtdEmTesteM += q.QtdEmTesteM
		quadroGrupo.QtdRecusado += q.QtdRecusado
		quadroGrupo.QtdBlocante += q.QtdBlocante

		grupoProjetos[q.Projeto.Grupo] = quadroGrupo
	}

	var keys projeto.Grupos
	for gp := range grupoProjetos {
		keys = append(keys, gp)
	}
	sort.Sort(keys)

	for _, grupo := range keys {
		buffer.WriteString(gerarQuadroTesteItem(grupo.ToString(), grupoProjetos[grupo]))
	}

	buffer.WriteString("</table>")

	return buffer.String()
}

func gerarQuadroTesteItem(produto string, quadro models.Quadro) string {
	var buffer bytes.Buffer

	buffer.WriteString("<tr>")
	buffer.WriteString("<td>" + produto + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalEmDesenvolvimento()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdAgMerge) + "</td>")

	if quadro.QtdAgTestePB > 0 {
		buffer.WriteString("<td style=\"background-color: red;\">" + strconv.Itoa(quadro.TotalAgTeste()) + " (" + strconv.Itoa(quadro.QtdAgTestePB) + ")" + "</td>")
	} else {
		buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalAgTeste()) + " (" + strconv.Itoa(quadro.QtdAgTestePB) + ")" + "</td>")
	}

	if quadro.QtdEmTestePB > 0 {
		buffer.WriteString("<td style=\"background-color: red;\">" + strconv.Itoa(quadro.TotalEmTeste()) + " (" + strconv.Itoa(quadro.QtdEmTestePB) + ")" + "</td>")
	} else {
		buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalEmTeste()) + " (" + strconv.Itoa(quadro.QtdEmTestePB) + ")" + "</td>")
	}
	buffer.WriteString("<td>" + strconv.Itoa(quadro.QtdRecusado) + "</td>")
	if quadro.QtdBlocante > 0 {
		buffer.WriteString("<td style=\"background-color: red;\">" + strconv.Itoa(quadro.QtdBlocante) + "</td>")
	} else {
		buffer.WriteString("<td>0</td>")
	}
	buffer.WriteString("</tr>")

	return buffer.String()
}
