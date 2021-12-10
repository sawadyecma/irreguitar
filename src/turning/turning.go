package turning

import (
	"reflect"

	"github.com/sawadyecma/irreguitar/sound"
)

var DefaultThreadCnt int = 6

type turning struct {
	threads map[ThreadNum]Thread
}

func NewTurning(threads map[ThreadNum]Thread) Turning {
	return turning{
		threads: threads,
	}
}

func (r turning) Threads() map[ThreadNum]Thread {
	s := reflect.ValueOf(r.threads)

	ret := make(map[ThreadNum]Thread, s.Len())

	for _, key := range s.MapKeys() {
		ret[key.Interface().(ThreadNum)] = s.MapIndex(key).Interface().(Thread)
	}

	return ret
}

func (r turning) ThreadNums() []ThreadNum {
	keys := reflect.ValueOf(r.threads).MapKeys()
	ret := make([]ThreadNum, len(r.threads))
	for i := range keys {
		ret[i] = keys[i].Interface().(ThreadNum)
	}
	return ret
}

func (r turning) Note(thnm ThreadNum, flet int) sound.Absnote {
	th, ok := r.threads[thnm]
	if !ok {
		panic("invalid thnm")
	}
	return th.OpenNote().Up(flet)
}
