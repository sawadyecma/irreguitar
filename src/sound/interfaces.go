package sound

import "fmt"

type Absnote interface {
	fmt.Stringer
	Up(halfToneDiff int) Absnote
	Diff(target Absnote) int
}

type Chord interface {
	Name() string
}
