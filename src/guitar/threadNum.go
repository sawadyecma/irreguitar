package guitar

import (
	"errors"

	"github.com/sawadyecma/irreguitar/sound"
)

func NewThreadNum(i int) (*threadNum, error) {
	if 1 <= i && i <= 6 {
		n := threadNum(i)
		return &n, nil
	}
	return nil, errors.New("invalid ThreadNum")
}

func (r threadNum) RegularOpenNote() sound.Absnote {
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
