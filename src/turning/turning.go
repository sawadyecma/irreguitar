package turning

import (
	"reflect"

	"github.com/sawadyecma/irreguitar/guitar"
	"github.com/sawadyecma/irreguitar/sound"
)

var DefaultThreadCnt int = 6

type turning struct {
	threads map[guitar.ThreadNum]guitar.Thread
}

func NewTurning(threads map[guitar.ThreadNum]guitar.Thread) Turning {
	return turning{
		threads: threads,
	}
}

func (r turning) Threads() map[guitar.ThreadNum]guitar.Thread {
	s := reflect.ValueOf(r.threads)

	ret := make(map[guitar.ThreadNum]guitar.Thread, s.Len())

	for _, key := range s.MapKeys() {
		ret[key.Interface().(guitar.ThreadNum)] = s.MapIndex(key).Interface().(guitar.Thread)
	}

	return ret
}

func (r turning) ThreadNums() []guitar.ThreadNum {
	keys := reflect.ValueOf(r.threads).MapKeys()
	ret := make([]guitar.ThreadNum, len(r.threads))
	for i := range keys {
		ret[i] = keys[i].Interface().(guitar.ThreadNum)
	}
	return ret
}

func (r turning) Note(thnm guitar.ThreadNum, flet int) sound.Absnote {
	th, ok := r.threads[thnm]
	if !ok {
		panic("invalid thnm")
	}
	return th.OpenNote().Up(flet)
}
