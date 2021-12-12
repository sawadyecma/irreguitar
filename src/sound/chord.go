package sound

import (
	"fmt"
	"reflect"
)

type chord struct {
	// not lowest note
	rootNote absnote
	notes    []absnote
	name     string
}

func NewChord(rootNote Absnote, notes []Absnote, name string) Chord {

	noteVals := make([]absnote, len(notes))
	for i := range notes {
		noteVals[i] = reflect.ValueOf(notes[i]).Interface().(absnote)
	}
	return chord{
		rootNote: reflect.ValueOf(rootNote).Interface().(absnote),
		notes:    noteVals,
		name:     name,
	}
}

func (c chord) Name() string {
	return c.name
}
