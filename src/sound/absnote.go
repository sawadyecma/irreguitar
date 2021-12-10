package sound

type absnote int

const (
	C1 absnote = iota
	Db1
	D1
	Eb1
	E1
	F1
	Gb1
	G1
	Ab1
	A1
	Bb1
	B1
	C2
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
	Gb4
	G4
	Ab4
	A4
	Bb4
	B4
	C5
	Db5
	D5
	Eb5
	E5
	F5
	Gb5
	G5
	Ab5
	A5
	Bb5
	B5
	C6
	Db6
	D6
	Eb6
	E6
	F6
	Gb6
	G6
	Ab6
	A6
	Bb6
	B6
)

var absnoteToString = map[absnote]string{
	C1:  "C1",
	Db1: "Db1",
	D1:  "D1",
	Eb1: "Eb1",
	E1:  "E1",
	F1:  "F1",
	Gb1: "Gb1",
	G1:  "G1",
	Ab1: "Ab1",
	A1:  "A1",
	Bb1: "Bb1",
	B1:  "B1",
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
	Gb4: "Gb4",
	G4:  "G4",
	Ab4: "Ab4",
	A4:  "A4",
	Bb4: "Bb4",
	B4:  "B4",
	C5:  "C5",
	Db5: "Db5",
	D5:  "D5",
	Eb5: "Eb5",
	E5:  "E5",
	F5:  "F5",
	Gb5: "Gb5",
	G5:  "G5",
	Ab5: "Ab5",
	A5:  "A5",
	Bb5: "Bb5",
	B5:  "B5",
	C6:  "C6",
	Db6: "Db6",
	D6:  "D6",
	Eb6: "Eb6",
	E6:  "E6",
	F6:  "F6",
	Gb6: "Gb6",
	G6:  "G6",
	Ab6: "Ab6",
	A6:  "A6",
	Bb6: "Bb6",
	B6:  "B6",
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
