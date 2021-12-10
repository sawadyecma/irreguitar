package sound

import "fmt"

type Absnote interface {
	fmt.Stringer
	Up(halfToneDiff int) Absnote
}

type Chord interface {
	Name() string
}
