package main

import (
	"fmt"

	"github.com/sawadyecma/irreguitar/guitar"
)

func main() {
	fmt.Println("[Start]Irreguitar!")

	printTurning()
}

func printTurning() {
	fmt.Println("===6 Strings Guitar Turning===")
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
