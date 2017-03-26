package acelerato

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"strconv"
	"bytes"
	"github.com/mbier/aceleratomo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/mbier/aceleratomo/mongo"
	"time"
)

const (
	Usuario string = "marlon.bier@mosistemas.com"
	Token string = "5cGDISvo1gM2Gi7tO7G+jA=="
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

func getDemandasSMOFrete() []models.Demanda {

	url := "https://mosistemas.acelerato.com/api/demandas?projetos=4&categorias=22"

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

func GerarQuadroTrack(mongoSession *mgo.Session) string {

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := GerarDadosQuadroTrack(mongoSession)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

//func ValidarGravacaoAutoTrack(mongoSession *mgo.Session) bool {
//	var date time.Time = time.Now()
//	return mongoSession.DB(mongo.AuthDatabase).C("projeto_dados").Find(bson.M{"tipo_projeto": models.TRACK, "data_geracao": {"$gte": "new Date(2017-03-12)"}}).Count() > 0
//}

func GerarDadosQuadroTrack(mongoSession *mgo.Session) (qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria int) {
	demandas := getDemandasTrack()

	testeFiltro := []int{10, 11}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria = gerarQuadro(demandas, testeFiltro)

	gravarProjetoDados(mongoSession, models.TRACK, qtdBacklogMelhoria, qtdBacklogProblema, qtdTesteMelhoria, qtdTesteProblema)

	return
}

func GerarQuadroAdm(mongoSession *mgo.Session) string {
	demandas := getDemandasAdm()

	testeFiltro := []int{10, 11}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	gravarProjetoDados(mongoSession, models.ADM, qtdBacklogMelhoria, qtdBacklogProblema, qtdTesteMelhoria, qtdTesteProblema)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroTMSWEB(mongoSession *mgo.Session) string {
	demandas := getDemandasTMSWEB()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	gravarProjetoDados(mongoSession, models.TMS_WEB, qtdBacklogMelhoria, qtdBacklogProblema, qtdTesteMelhoria, qtdTesteProblema)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroSMONET(mongoSession *mgo.Session) string {
	demandas := getDemandasSMONET()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	gravarProjetoDados(mongoSession, models.SMO_NET, qtdBacklogMelhoria, qtdBacklogProblema, qtdTesteMelhoria, qtdTesteProblema)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroSMOWEB(mongoSession *mgo.Session) string {
	demandas := getDemandasSMOWEB()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	gravarProjetoDados(mongoSession, models.SMO_WEB, qtdBacklogMelhoria, qtdBacklogProblema, qtdTesteMelhoria, qtdTesteProblema)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroSMOCTE(mongoSession *mgo.Session) string {
	demandas := getDemandasSMOCTE()

	testeFiltro := []int{19, 20}

	qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria := gerarQuadro(demandas, testeFiltro)

	gravarProjetoDados(mongoSession, models.SMO_CTE, qtdBacklogMelhoria, qtdBacklogProblema, qtdTesteMelhoria, qtdTesteProblema)

	return gerarQuadroString(qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria)
}

func GerarQuadroGeral() string {
	demandasSmocte := getDemandasSMOCTE()
	demandasTmsweb := getDemandasTMSWEB()
	demandasSmonet := getDemandasSMONET()
	demandasTrack := getDemandasTrack()
	demandasAdm := getDemandasAdm()
	demandasLogrev := getDemandasLogrev()

	testeFiltroTrack := []int{10, 11}
	testeFiltroFlex := []int{10, 11}

	qtdBacklogProblemaTrack, qtdBacklogMelhoriaTrack, qtdTesteProblemaTrack, qtdTesteMelhoriaTrack := gerarQuadro(demandasTrack, testeFiltroTrack)
	qtdBacklogProblemaAdm, qtdBacklogMelhoriaAdm, qtdTesteProblemaAdm, qtdTesteMelhoriaAdm := gerarQuadro(demandasAdm, testeFiltroTrack)
	qtdBacklogProblemaTmsweb, qtdBacklogMelhoriaTmsweb, qtdTesteProblemaTmsweb, qtdTesteMelhoriaTmsweb := gerarQuadro(demandasTmsweb, testeFiltroFlex)
	qtdBacklogProblemaSmonet, qtdBacklogMelhoriaSmonet, qtdTesteProblemaSmonet, qtdTesteMelhoriaSmonet := gerarQuadro(demandasSmonet, testeFiltroFlex)
	qtdBacklogProblemaSmocte, qtdBacklogMelhoriaSmocte, qtdTesteProblemaSmocte, qtdTesteMelhoriaSmocte := gerarQuadro(demandasSmocte, testeFiltroFlex)
	qtdBacklogProblemaLogrev, qtdBacklogMelhoriaLogrev, qtdTesteProblemaLogrev, qtdTesteMelhoriaLogrev := gerarQuadro(demandasLogrev, testeFiltroFlex)

	qtdBacklogProblema := 0
	qtdBacklogMelhoria := 0
	qtdTesteProblema := 0
	qtdTesteMelhoria := 0

	qtdBacklogProblemaGeral := qtdBacklogProblemaTrack + qtdBacklogProblemaTmsweb + qtdBacklogProblemaSmonet + qtdBacklogProblemaAdm + qtdBacklogProblemaSmocte + qtdBacklogProblemaLogrev
	qtdBacklogMelhoriaGeral := qtdBacklogMelhoriaTrack + qtdBacklogMelhoriaTmsweb + qtdBacklogMelhoriaSmonet + qtdBacklogMelhoriaAdm + qtdBacklogMelhoriaSmocte + qtdBacklogMelhoriaLogrev
	qtdTesteProblemaGeral := qtdTesteProblemaTrack + qtdTesteProblemaTmsweb + qtdTesteProblemaSmonet + qtdTesteProblemaAdm + qtdTesteProblemaSmocte + qtdTesteProblemaLogrev
	qtdTesteMelhoriaGeral := qtdTesteMelhoriaTrack + qtdTesteMelhoriaTmsweb + qtdTesteMelhoriaSmonet + qtdTesteMelhoriaAdm + qtdTesteMelhoriaSmocte + qtdTesteMelhoriaLogrev

	var buffer bytes.Buffer

	buffer.WriteString("<table style=\"width:100%\">")
	buffer.WriteString("<tr><th>Produto</th><th>Melhoria</th><th>Problema</th><th>AG. Teste</th><th>Total</th><th>&#37; Melhoria</th><th>&#37; Problema</th></tr>")

	buffer.WriteString(gerarQuadroGeralItem("SMOCTE", qtdBacklogProblemaSmocte, qtdBacklogMelhoriaSmocte, qtdTesteProblemaSmocte, qtdTesteMelhoriaSmocte))
	buffer.WriteString(gerarQuadroGeralItem("SMOTMS", qtdBacklogProblemaTmsweb, qtdBacklogMelhoriaTmsweb, qtdTesteProblemaTmsweb, qtdTesteMelhoriaTmsweb))
	buffer.WriteString(gerarQuadroGeralItem("ADM", qtdBacklogProblemaAdm, qtdBacklogMelhoriaAdm, qtdTesteProblemaAdm, qtdTesteMelhoriaAdm))
	buffer.WriteString(gerarQuadroGeralItem("SMO-FRETE", qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria))
	buffer.WriteString(gerarQuadroGeralItem("LOGREV", qtdBacklogProblemaLogrev, qtdBacklogMelhoriaLogrev, qtdTesteProblemaLogrev, qtdTesteMelhoriaLogrev))
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

func gerarQuadro(demandas []models.Demanda, testeFilter[]int) (qtdBacklogProblema, qtdBacklogMelhoria, qtdTesteProblema, qtdTesteMelhoria int) {
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
func gravarProjetoDados(mongoSession *mgo.Session, tipoProjeto models.TipoProjeto, qtdBacklogMelhoria, qtdBacklogProblema, qtdTesteMelhoria, qtdTesteProblema int) {
	if mongoSession == nil{
		return
	}

	projeto := models.Projeto{}

	projeto.ID = bson.NewObjectId()
	projeto.Data = time.Now()
	projeto.TipoProjeto = tipoProjeto
	projeto.QtdBacklogMelhoria = qtdBacklogMelhoria
	projeto.QtdBacklogProblema = qtdBacklogProblema
	projeto.QtdTesteMelhoria = qtdTesteMelhoria
	projeto.QtdTesteProblema = qtdTesteProblema

	err := mongoSession.DB(mongo.AuthDatabase).C("projeto_dados").Insert(projeto)
	if err != nil {
		log.Println(err)
	}
}