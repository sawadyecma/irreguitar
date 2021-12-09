package sound

type absnote int

const (
	E2 absnote = iota
	A2
	D3
	G3
	B3
	E4
)

func (r absnote) String() string {
	switch r {
	case E2:
		return "E2"
	case A2:
		return "A2"
	case D3:
		return "D3"
	case G3:
		return "G3"
	case B3:
		return "B3"
	case E4:
		return "E4"
	default:
		panic("invalid absnote")
	}
}
