package pontocardeal

type Direcao int

const (
	Norte Direcao = iota
	Sul
	Leste
	Oeste
)

var value = []string{"Norte", "Sul", "Leste", "Oeste"}

func (p Direcao) String() string {
	return value[p]
}
