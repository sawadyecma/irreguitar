package turning

import (
	"github.com/sawadyecma/irreguitar/sound"
)

type Turning interface {
	Threads() map[ThreadNum]Thread
	ThreadNums() []ThreadNum
	Note(thnm ThreadNum, flet int) sound.Absnote
}

type ThreadNum int

type Thread interface {
	OpenNote() sound.Absnote
}
