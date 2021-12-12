package sound

import "fmt"

type Absnote interface {
	fmt.Stringer
	Up(halfToneDiff int) Absnote
	Diff(target Absnote) int
	Name() string
}

type Chord interface {
	Name() string
}

type ChordGenerator interface {
	Generate(root Absnote, notes []Absnote) Chord
}

type ChordParser interface {
	Parse(root Absnote, notes []Absnote) string
}
