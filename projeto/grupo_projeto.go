package projeto

// Grupo representa o grupo de projeto
type Grupo int

const (
	ADM Grupo = iota
	FRETE
	DELPHI
	FLEX
)

// ToString exibe o nome
func (g Grupo) ToString() string {
	switch g {
	case 0:
		return "ADM"
	case 1:
		return "FRETE"
	case 2:
		return "DELPHI"
	case 3:
		return "FLEX"
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
