package sound

type absnote int

const (
	C2 absnote = iota
	Db2
	D2
	Eb2
	E2
	F2
	Gb2
	G2
	Ab2
	A2
	Bb2
	B2
	C3
	Db3
	D3
	Eb3
	E3
	F3
	Gb3
	G3
	Ab3
	A3
	Bb3
	B3
	C4
	Db4
	D4
	Eb4
	E4
	F4
)

var absnoteToString = map[absnote]string{
	C2:  "C2",
	Db2: "Db2",
	D2:  "D2",
	Eb2: "Eb2",
	E2:  "E2",
	F2:  "F2",
	Gb2: "Gb2",
	G2:  "G2",
	Ab2: "Ab2",
	A2:  "A2",
	Bb2: "Bb2",
	B2:  "B2",
	C3:  "C3",
	Db3: "Db3",
	D3:  "D3",
	Eb3: "Eb3",
	E3:  "E3",
	F3:  "F3",
	Gb3: "Gb3",
	G3:  "G3",
	Ab3: "Ab3",
	A3:  "A3",
	Bb3: "Bb3",
	B3:  "B3",
	C4:  "C4",
	Db4: "Db4",
	D4:  "D4",
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

func (r absnote) Up(halfToneDiff int) Absnote {
	return absnote(int(r) + halfToneDiff)
}
