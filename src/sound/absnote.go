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

var absnoteToString = map[absnote]string{
	Eb2: "Eb2",
	E2:  "E2",
	F2:  "F2",
	Ab2: "Ab2",
	A2:  "A2",
	Bb2: "Bb2",
	Db3: "Db3",
	D3:  "D3",
	Eb3: "Eb3",
	Gb3: "Gb3",
	G3:  "G3",
	Ab3: "Ab3",
	Bb3: "Bb3",
	B3:  "B3",
	C4:  "C4",
	Eb4: "Eb4",
	E4:  "E4",
	F4:  "F4",
}

func (r absnote) String() string {
	v, ok := absnoteToString[r]
	if !ok {
		panic("invalid absnote")
	}
	return v
}

func (r absnote) TurnUp(halfToneDiff int) Absnote {
	return absnote(int(r) + halfToneDiff)
}
