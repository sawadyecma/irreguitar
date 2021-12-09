package main

import (
	"fmt"

	"github.com/sawadyecma/irreguitar/guitar"
)

func main() {
	fmt.Println("Hello From Irreguitar!")
	tn := guitar.NewTurning()
	threads := tn.Threads()
	for i := range threads {
		fmt.Printf(
			"number: %d, note: %s\n",
			threads[i].Num(),
			threads[i].OpenNote().String(),
		)
	}
}
