package play

import (
	"github.com/sawadyecma/irreguitar/sound"
)

type Play interface {
	Chord(rootPress Press, presses []Press) sound.Chord
}
