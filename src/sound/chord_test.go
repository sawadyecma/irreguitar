package sound

import (
	"fmt"
)

func Example_chord_Name() {
	fmt.Println(
		NewChord(C3, []Absnote{E3, G3}).Name(),
	)

	fmt.Println(
		NewChord(E3, []Absnote{E3, G3}).Name(),
	)

	fmt.Println(
		NewChord(C3, []Absnote{E2, G3}).Name(),
	)

	// Output:
	// C3
	// E3
	// C3/E2
}
