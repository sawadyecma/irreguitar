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

	fmt.Println("===6 Strings Guitar Turning2===")
	threads := make(map[guitar.ThreadNum]guitar.Thread, guitar.DefaultThreadCnt)
	for i := 1; i <= guitar.DefaultThreadCnt; i++ {
		threadNum, err := guitar.NewThreadNum(i)
		if err != nil {
			panic(err)
		}
		threads[*threadNum] = guitar.NewThread(
			threadNum.RegularOpenNote(),
		)
	}

	tn2 := guitar.NewTurning(threads)
	printTurning(tn2)
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
