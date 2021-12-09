package guitar

import (
	"github.com/sawadyecma/irreguitar/sound"
)

type Turning interface {
	Threads() map[ThreadNum]Thread
	ThreadNums() []ThreadNum
}

type ThreadNum int

type Thread interface {
	OpenNote() sound.Absnote
}
