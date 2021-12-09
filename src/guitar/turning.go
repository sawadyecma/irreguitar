package guitar

import (
	"fmt"
	"reflect"
)

var gtThreadCnt int = 6

type turning struct {
	threads []thread
}

func NewTurning() Turning {
	threads := make([]thread, gtThreadCnt)
	for i := 1; i <= gtThreadCnt; i++ {
		threadNum, err := NewThreadNum(i)
		if err != nil {
			fmt.Printf("err: %s\n", err)
			return nil
		}
		threads[i-1] = NewThread(*threadNum)
	}

	return turning{
		threads: threads,
	}
}

func (r turning) Threads() []Thread {
	s := reflect.ValueOf(r.threads)

	ret := make([]Thread, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface().(Thread)
	}

	return ret
}
