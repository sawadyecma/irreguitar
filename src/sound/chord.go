package sound

import (
	"reflect"
)

const (
	major = iota
	minor
	majerSeventh
	minorSeventh
	seventh
	nineth
	addNine
)

var codeTypeToString = map[int]string{
	major:        "",
	minor:        "m",
	majerSeventh: "M7",
	minorSeventh: "m7",
	seventh:      "7",
	nineth:       "9",
	addNine:      "add9",
}

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
