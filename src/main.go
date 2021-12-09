package main

import (
	"fmt"

	"github.com/sawadyecma/irreguitar/guitar"
)

func main() {
	fmt.Println("[Start]Irreguitar!")

	fmt.Println("===6 Strings Guitar Regular Turning===")
	tn := guitar.NewRegularTurning()
	printTurning(tn)

	fmt.Println("===6 Strings Guitar HalfDown Turning===")
	halfDownTurning := newTurning(
		map[int]int{6: -1, 5: -1, 4: -1, 3: -1, 2: -1, 1: -1},
	)
	printTurning(halfDownTurning)

	fmt.Println("===6 Strings Guitar HalfUp Turning===")
	halfUpTurning := newTurning(
		map[int]int{6: +1, 5: +1, 4: +1, 3: +1, 2: +1, 1: +1},
	)
	printTurning(halfUpTurning)
}

func newTurning(turn map[int]int) guitar.Turning {
	threads := make(map[guitar.ThreadNum]guitar.Thread, guitar.DefaultThreadCnt)

	for i := 1; i <= guitar.DefaultThreadCnt; i++ {
		threadNum, err := guitar.NewThreadNum(i)
		if err != nil {
			panic(err)
		}
		threads[*threadNum] = guitar.NewThread(
			threadNum.RegularOpenNote().TurnUp(turn[i]),
		)
	}

	return guitar.NewTurning(threads)
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
