package guitar

import (
	"reflect"
)

var gtThreadCnt int = 6

type turning struct {
	threads map[threadNum]Thread
}

type turnConfig struct {
}

func NewTurning() Turning {
	threads := make(map[threadNum]Thread, gtThreadCnt)
	for i := 1; i <= gtThreadCnt; i++ {
		threadNum, err := NewThreadNum(i)
		if err != nil {
			panic(err)
		}
		threads[*threadNum] = NewThread(
			threadNum.RegularOpenNote(),
		)
	}

	return turning{
		threads: threads,
	}
}

func (r turning) Threads() map[threadNum]Thread {
	s := reflect.ValueOf(r.threads)

	ret := make(map[threadNum]Thread, s.Len())

	for _, key := range s.MapKeys() {
		ret[key.Interface().(threadNum)] = s.MapIndex(key).Interface().(Thread)
	}

	return ret
}

func (r turning) ThreadNums() []threadNum {
	keys := reflect.ValueOf(r.threads).MapKeys()
	ret := make([]threadNum, len(r.threads))
	for i := range keys {
		ret[i] = keys[i].Interface().(threadNum)
	}
	return ret
}
