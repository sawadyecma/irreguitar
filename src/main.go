package main

import (
	"fmt"

	"github.com/sawadyecma/irreguitar/guitar"
)

func main() {
	fmt.Println("[Start]Irreguitar!")
	threadCnt := 6

	tn := guitar.NewTurning(threadCnt)
	threads := tn.Threads()
	for i := range threads {
		fmt.Printf(
			"number: %d, note: %s\n",
			threads[i].Num(),
			threads[i].OpenNote().String(),
		)
	}
}
