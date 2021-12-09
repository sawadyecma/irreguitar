package guitar

import (
	"github.com/sawadyecma/irreguitar/sound"
)

type Turning interface {
	Threads() map[threadNum]Thread
	ThreadNums() []threadNum
}

type threadNum int

type Thread interface {
	OpenNote() sound.Absnote
}
