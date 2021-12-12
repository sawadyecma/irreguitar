package sound

import (
	"fmt"
)

type mockParser struct{}

func (r mockParser) Parse(root Absnote, notes []Absnote) string {
	return "Test"
}

func Example_chordGenerator_Generate() {

	g := NewChordGenerator(mockParser{})

	fmt.Println(
		g.Generate(C3, []Absnote{E3, G3}).Name(),
	)

	// Output:
	// Test
}
