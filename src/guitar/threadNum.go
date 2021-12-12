package guitar

import (
	"github.com/sawadyecma/irreguitar/sound"
)

func NewThreadNum(i int) *ThreadNum {
	if 1 <= i && i <= 6 {
		n := ThreadNum(i)
		return &n
	}
	panic("invalid ThreadNum")
}

func (r ThreadNum) RegularOpenNote() sound.Absnote {
	switch r {
	case 6:
		return sound.E2
	case 5:
		return sound.A2
	case 4:
		return sound.D3
	case 3:
		return sound.G3
	case 2:
		return sound.B3
	case 1:
		return sound.E4
	default:
		panic("invalid thread num")
	}
}
