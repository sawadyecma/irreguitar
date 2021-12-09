package guitar

import (
	"fmt"
	"reflect"
)

var gtThreadCnt int = 6

type turning struct {
	threads map[threadNum]thread
}

func NewTurning() Turning {
	threads := make(map[threadNum]thread, gtThreadCnt)
	for i := 1; i <= gtThreadCnt; i++ {
		threadNum, err := NewThreadNum(i)
		if err != nil {
			fmt.Printf("err: %s\n", err)
			return nil
		}
		threads[*threadNum] = NewThread(threadNum.Absnote())
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
