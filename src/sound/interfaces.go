package sound

import "fmt"

type Absnote interface {
	fmt.Stringer
	TurnUp(halfToneDiff int) Absnote
}
