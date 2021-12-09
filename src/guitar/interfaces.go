package guitar

import (
	"github.com/sawadyecma/irreguitar/sound"
)

type Turning interface {
	Threads() []Thread
}

type threadNum int

type Thread interface {
	Number() threadNum
	OpenNote() sound.Note
}
