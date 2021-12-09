package main

import (
	"fmt"

	"github.com/sawadyecma/irreguitar/guitar"
)

func main() {
	fmt.Println("[Start]Irreguitar!")

	fmt.Println("===6 Strings Guitar Turning===")
	tn := guitar.NewTurning()
	printTurning(tn)
}

func printTurning(tn guitar.Turning) {
	threads := tn.Threads()
	nums := tn.ThreadNums()
	fmt.Printf("threadNums: %v\n", nums)

	for _, num := range nums {
		fmt.Printf(
			"number: %d, note: %s\n",
			num,
			threads[num].OpenNote().String(),
		)
	}
}
