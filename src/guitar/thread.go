package guitar

import (
	"errors"

	"github.com/sawadyecma/irreguitar/sound"
)

func NewThread() Thread {
	return thread{}
}

type thread struct{}

func (thread) Number() threadNum {
	// TODO フィールドから取得
	tn, _ := NewThreadNum(1)
	return *tn
}

func (thread) OpenNote() sound.Note {
	return nil
}

func NewThreadNum(i int) (*threadNum, error) {
	if 1 <= i && i <= 6 {
		n := threadNum(i)
		return &n, nil
	}
	return nil, errors.New("invalid ThreadNum")
}
