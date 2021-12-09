package sound

type absnote int

const (
	Ef2 absnote = iota
	E2
	Af2
	A2
	Df3
	D3
	Gf3
	G3
	Bf3
	B3
	Ef4
	E4
)

var Absnotes = []absnote{
	Ef2, E2, A2, D3, G3, B3, E4,
}

func (r absnote) String() string {
	switch r {
	case Ef2:
		return "Ef2"
	case E2:
		return "E2"
	case Af2:
		return "Af2"
	case A2:
		return "A2"
	case Df3:
		return "Df3"
	case D3:
		return "D3"
	case Gf3:
		return "Gf3"
	case G3:
		return "G3"
	case Bf3:
		return "Bf3"
	case B3:
		return "B3"
	case Ef4:
		return "Ef4"
	case E4:
		return "E4"
	default:
		panic("invalid absnote")
	}
}

func (r absnote) Turn(halfTone int) Absnote {
	a := int(r)
	return absnote(a + halfTone)
}
