package guitar

import (
	"errors"

	"github.com/sawadyecma/irreguitar/sound"
)

func NewThread(n threadNum) thread {
	return thread{
		threadNum: n,
	}
}

type thread struct {
	threadNum threadNum
}

func (r thread) Num() threadNum {
	return r.threadNum
}

func (r thread) OpenNote() sound.Note {
	if r.threadNum == 6 {
		return sound.NewNote()
	}
	return sound.NewNote()
}

func NewThreadNum(i int) (*threadNum, error) {
	if 1 <= i && i <= 6 {
		n := threadNum(i)
		return &n, nil
	}
	return nil, errors.New("invalid ThreadNum")
}
