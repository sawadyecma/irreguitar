package turning

import (
	"github.com/sawadyecma/irreguitar/guitar"
	"github.com/sawadyecma/irreguitar/sound"
)

type Turning interface {
	Threads() map[guitar.ThreadNum]guitar.Thread
	ThreadNums() []guitar.ThreadNum
	Note(thnm guitar.ThreadNum, flet int) sound.Absnote
}
