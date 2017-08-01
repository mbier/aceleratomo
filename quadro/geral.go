package quadro

import (
	"bytes"
	"sort"
	"strconv"
	"sync"

	"github.com/mbier/aceleratomo/acelerato"
	"github.com/mbier/aceleratomo/projeto"
)

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

	ch := make(chan acelerato.ProjetoDado, len(projetos))
	wg := sync.WaitGroup{}

	for _, p := range projetos {
		wg.Add(1)
		go func(projeto projeto.Projeto) {
			quadro := acelerato.GerarDadosQuadro(&projeto)
			quadro.Projeto = projeto

			defer wg.Done()

			ch <- quadro
		}(p)
	}

	wg.Wait()
	close(ch)
	var quadros acelerato.ProjetosDado

	for n := range ch {
		quadros = append(quadros, n)
	}

	sort.Sort(quadros)

	for _, q := range quadros {

		buffer.WriteString(gerarQuadroGeralItem(q.Projeto.Nome, &q))

		qtdBacklogProblemaGeral += q.TotalBacklogP()
		qtdBacklogMelhoriaGeral += q.TotalBacklogM()
		qtdTesteProblemaGeral += q.TotalTesteP()
		qtdTesteMelhoriaGeral += q.TotalTesteM()
		qtdAgMergeGeral += q.TotalAgMerge()
		qtdImpedimentoGeral += q.QtdImpedimento
	}

	quadroGeral := acelerato.NewProjetoDado()
	quadroGeral.Aprovado.QtdMelhoria = qtdBacklogMelhoriaGeral
	quadroGeral.Aprovado.QtdProblema = qtdBacklogProblemaGeral
	quadroGeral.EmTeste.QtdProblema = qtdTesteProblemaGeral
	quadroGeral.EmTeste.QtdMelhoria = qtdTesteMelhoriaGeral
	quadroGeral.AgMerge.QtdMelhoria = qtdAgMergeGeral
	quadroGeral.QtdImpedimento = qtdImpedimentoGeral

	buffer.WriteString(gerarQuadroGeralItem("Total", &quadroGeral))

	buffer.WriteString("</table>")

	return buffer.String()
}

func gerarQuadroGeralItem(produto string, quadro *acelerato.ProjetoDado) string {
	var buffer bytes.Buffer

	buffer.WriteString("<tr>")
	buffer.WriteString("<td>" + produto + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogM()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalBacklogP()) + "</td>")
	buffer.WriteString("<td>" + strconv.Itoa(quadro.TotalAgMerge()) + "</td>")
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
