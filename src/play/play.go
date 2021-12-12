package play

import (
	"github.com/sawadyecma/irreguitar/guitar"
	"github.com/sawadyecma/irreguitar/sound"
	"github.com/sawadyecma/irreguitar/turning"
)

type play struct {
	turning turning.Turning
	cg      sound.ChordGenerator
}

type Press struct {
	threadNum guitar.ThreadNum
	flet      int
}

func NewPlay(
	turning turning.Turning,
	cg sound.ChordGenerator,
) Play {
	return play{
		turning: turning,
		cg:      cg,
	}
}

func (p play) Chord(rootPress Press, presses []Press) sound.Chord {
	threads := p.turning.Threads()
	rootNote := threads[rootPress.threadNum].Note(rootPress.flet)

	notes := make([]sound.Absnote, len(presses))
	for i := range presses {
		press := presses[i]
		notes[i] = threads[press.threadNum].Note(press.flet)
	}

	return p.cg.Generate(
		rootNote,
		notes,
	)
}
