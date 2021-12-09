package guitar

import (
	"github.com/sawadyecma/irreguitar/sound"
)

type Turning interface {
	Threads() []Thread
}

type threadNum int

type Thread interface {
	Num() threadNum
	OpenNote() sound.Note
}
