package sound

type absnote int

const (
	Eb2 absnote = iota
	E2
	F2
	Ab2
	A2
	Bb2
	Db3
	D3
	Eb3
	Gb3
	G3
	Ab3
	Bb3
	B3
	C4
	Eb4
	E4
	F4
)

func (r absnote) String() string {
	switch r {
	case Eb2:
		return "Eb2"
	case E2:
		return "E2"
	case F2:
		return "F2"
	case Ab2:
		return "Ab2"
	case A2:
		return "A2"
	case Bb2:
		return "Bb2"
	case Db3:
		return "Db3"
	case D3:
		return "D3"
	case Eb3:
		return "Eb3"
	case Gb3:
		return "Gb3"
	case G3:
		return "G3"
	case Ab3:
		return "Ab3"
	case Bb3:
		return "Bb3"
	case B3:
		return "B3"
	case C4:
		return "C4"
	case Eb4:
		return "Eb4"
	case E4:
		return "E4"
	case F4:
		return "F4"
	default:
		panic("invalid absnote")
	}
}

func (r absnote) TurnUp(halfToneDiff int) Absnote {
	return absnote(int(r) + halfToneDiff)
}
