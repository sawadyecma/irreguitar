package guitar

import (
	"github.com/sawadyecma/irreguitar/sound"
)

func NewThread(openNote sound.Absnote) Thread {
	return thread{
		openNote: openNote,
	}
}

type thread struct {
	openNote sound.Absnote
}

func (r thread) OpenNote() sound.Absnote {
	return r.openNote
}
