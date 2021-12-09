package sound

import "fmt"

type Absnote interface {
	fmt.Stringer
	Turn(halfTone int) Absnote
}
