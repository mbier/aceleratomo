package projeto

import "encoding/json"

// Grupo representa o grupo de projeto
type Grupo int

const (
	PORTAL Grupo = iota
	DELPHI
	FLEX
	ADM
	FRETE
	TMS
)

// ToString exibe o nome
func (g Grupo) ToString() string {
	switch g {
	case 0:
		return "PORTAL"
	case 1:
		return "DELPHI"
	case 2:
		return "FLEX"
	case 3:
		return "ADM"
	case 4:
		return "FRETE"
	case 5:
		return "TMS"
	}

	return ""
}

// Grupos lista de grupo
type Grupos []Grupo

func (q Grupos) Len() int {
	return len(q)
}

func (q Grupos) Less(i, j int) bool {
	return q[i].ToString() < q[j].ToString()
}

func (q Grupos) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (g *Grupo) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.ToString())
}
