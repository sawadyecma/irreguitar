package sound

import (
	"fmt"
	"reflect"
)

type chord struct {
	// not lowest note
	rootNote absnote
	notes    []absnote
}

func NewChord(rootNote Absnote, notes []Absnote) Chord {

	noteVals := make([]absnote, len(notes))
	for i := range notes {
		noteVals[i] = reflect.ValueOf(notes[i]).Interface().(absnote)
	}
	return chord{
		rootNote: reflect.ValueOf(rootNote).Interface().(absnote),
		notes:    noteVals,
	}
}

func (c chord) Name() string {
	// example: C Cm CM7 Cm7 C7 C9 Cadd9 Cm9 C11 C13 C7/E
	return fmt.Sprintf("%s", c.rootNote)
}
