package guitar

import (
	"github.com/sawadyecma/irreguitar/sound"
)

func NewThread(an sound.Absnote) thread {
	return thread{
		openNote: an,
	}
}

type thread struct {
	openNote sound.Absnote
}

func (r thread) OpenNote() sound.Absnote {
	return r.openNote
}
