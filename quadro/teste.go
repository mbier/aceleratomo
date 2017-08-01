package quadro

import (
	"bytes"
	"sort"
	"strconv"
	"sync"

	"github.com/mbier/aceleratomo/acelerato"
	"github.com/mbier/aceleratomo/projeto"
)

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
	buffer.WriteString("<tr><th>Equipe</th><th>Desenv</th><th>AG. Merge</th><th>AG. Teste</th><th>Em Teste</th></tr>")

	projetos := projeto.GetProjetos()

	out := make(chan acelerato.ProjetoDado, len(projetos))
	wg := sync.WaitGroup{}

	for _, p := range projetos {
		wg.Add(1)
		go func(projeto projeto.Projeto) {
			quadro := acelerato.GerarDadosQuadro(&projeto)
			quadro.Projeto = projeto

			defer wg.Done()

			out <- quadro
		}(p)
	}

	wg.Wait()
	close(out)

	var quadros acelerato.ProjetosDado

	for n := range out {
		quadros = append(quadros, n)
	}

	grupoProjetos := make(map[projeto.Grupo]acelerato.ProjetoDado)

	for _, q := range quadros {

		quadroGrupo := grupoProjetos[q.Projeto.Grupo]

		quadroGrupo.Projeto.Nome = q.Projeto.Grupo.ToString()

		quadroGrupo.Aprovado.Merge(q.Aprovado)
		quadroGrupo.EmDesenvolvimento.Merge(q.EmDesenvolvimento)
		quadroGrupo.AgMerge.Merge(q.AgMerge)
		quadroGrupo.AgTeste.Merge(q.AgTeste)
		quadroGrupo.EmTeste.Merge(q.EmTeste)
		quadroGrupo.QtdImpedimento += q.QtdImpedimento

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

	buffer.WriteString("<div style=\"text-align: center;\"> B = Blocante / R = Recusa</div>")

	return buffer.String()
}

func gerarQuadroTesteItem(produto string, quadro acelerato.ProjetoDado) string {
	var buffer bytes.Buffer

	buffer.WriteString("<tr>")
	buffer.WriteString("<td>" + produto + "</td>")
	if quadro.EmDesenvolvimento.QtdBlocante > 0 {
		buffer.WriteString("<td style=\"background-color: red; color: white;\">" + strconv.Itoa(quadro.EmDesenvolvimento.Total()) + " (B:" + strconv.Itoa(quadro.EmDesenvolvimento.QtdBlocante) + " R:" + strconv.Itoa(quadro.EmDesenvolvimento.QtdRecusado) + ")" + "</td>")
	} else {
		buffer.WriteString("<td>" + strconv.Itoa(quadro.EmDesenvolvimento.Total()) + " (B:" + strconv.Itoa(quadro.EmDesenvolvimento.QtdBlocante) + " R:" + strconv.Itoa(quadro.EmDesenvolvimento.QtdRecusado) + ")" + "</td>")
	}
	if quadro.AgMerge.QtdBlocante > 0 {
		buffer.WriteString("<td style=\"background-color: red; color: white;\">" + strconv.Itoa(quadro.AgMerge.Total()) + " (B:" + strconv.Itoa(quadro.AgMerge.QtdBlocante) + " R:" + strconv.Itoa(quadro.AgMerge.QtdRecusado) + ")" + "</td>")
	} else {
		buffer.WriteString("<td>" + strconv.Itoa(quadro.AgMerge.Total()) + " (B:" + strconv.Itoa(quadro.AgMerge.QtdBlocante) + " R:" + strconv.Itoa(quadro.AgMerge.QtdRecusado) + ")" + "</td>")
	}
	if quadro.AgTeste.QtdBlocante > 0 {
		buffer.WriteString("<td style=\"background-color: red; color: white;\">" + strconv.Itoa(quadro.AgTeste.Total()) + " (B:" + strconv.Itoa(quadro.AgTeste.QtdBlocante) + " R:" + strconv.Itoa(quadro.AgTeste.QtdRecusado) + ")" + "</td>")
	} else {
		buffer.WriteString("<td>" + strconv.Itoa(quadro.AgTeste.Total()) + " (B:" + strconv.Itoa(quadro.AgTeste.QtdBlocante) + " R:" + strconv.Itoa(quadro.AgTeste.QtdRecusado) + ")" + "</td>")
	}
	if quadro.EmTeste.QtdBlocante > 0 {
		buffer.WriteString("<td style=\"background-color: red; color: white;\">" + strconv.Itoa(quadro.EmTeste.Total()) + " (B:" + strconv.Itoa(quadro.EmTeste.QtdBlocante) + " R:" + strconv.Itoa(quadro.EmTeste.QtdRecusado) + ")" + "</td>")
	} else {
		buffer.WriteString("<td>" + strconv.Itoa(quadro.EmTeste.Total()) + " (B:" + strconv.Itoa(quadro.EmTeste.QtdBlocante) + " R:" + strconv.Itoa(quadro.EmTeste.QtdRecusado) + ")" + "</td>")
	}
	buffer.WriteString("</tr>")

	return buffer.String()
}
