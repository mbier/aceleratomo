package quadro

import (
	"bytes"
	"strconv"

	"github.com/mbier/aceleratomo/acelerato"
	"github.com/mbier/aceleratomo/projeto"
)

// GerarQuadro com as informacoes do projeto
func GerarQuadro(projeto *projeto.Projeto) string {

	quadro := acelerato.GerarDadosQuadro(projeto)

	return gerarQuadroString(quadro)
}

func gerarQuadroString(quadro acelerato.ProjetoDado) string {
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

	if quadro.TotalAgMerge() > 0 {
		buffer.WriteString("<tr><td>Ag Merge</td><td>" + strconv.Itoa(quadro.TotalAgMerge()) + "</td>")
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
