package guitar

import "github.com/sawadyecma/irreguitar/sound"

type ThreadNum int

type Thread interface {
	OpenNote() sound.Absnote
	Note(flet int) sound.Absnote
}
