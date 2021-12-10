package main

import (
	"fmt"

	"github.com/sawadyecma/irreguitar/guitar"
)

func main() {
	fmt.Println("[Start]Irreguitar!")

	turnings := map[string]map[int]int{
		"regular":   {6: +0, 5: +0, 4: +0, 3: +0, 2: +0, 1: +0},
		"half down": {6: -1, 5: -1, 4: -1, 3: -1, 2: -1, 1: -1},
		"half up":   {6: +1, 5: +1, 4: +1, 3: +1, 2: +1, 1: +1},
		"dropD":     {6: -2, 5: +0, 4: +0, 3: +0, 2: +0, 1: +0},
		"DADGAD":    {6: -2, 5: +0, 4: +0, 3: +0, 2: -2, 1: -2},
		"DAEAC#E":   {6: -2, 5: +0, 4: +2, 3: +2, 2: +2, 1: +0},
		"FACGCE":    {6: +1, 5: +0, 4: -2, 3: +0, 2: +1, 1: +0},
		"DADAC#E":   {6: -2, 5: +0, 4: +0, 3: +2, 2: +2, 1: +0},
	}

	for k, v := range turnings {
		fmt.Printf("=== Turning: %s===\n", k)
		tn := newTurning(v)
		printTurning(tn)
	}

	th := newTurning(turnings["regular"])
	thnm, err := guitar.NewThreadNum(3)
	if err != nil {
		panic(err)
	}

	n := th.Note(*thnm, 3)
	fmt.Println(n)
	n = th.Note(*thnm, 2)
	fmt.Println(n)
	n = th.Note(*thnm, 1)
	fmt.Println(n)
	n = th.Note(*thnm, 0)
	fmt.Println(n)
}

func newTurning(turn map[int]int) guitar.Turning {
	threads := make(map[guitar.ThreadNum]guitar.Thread, guitar.DefaultThreadCnt)

	for i := 1; i <= guitar.DefaultThreadCnt; i++ {
		threadNum, err := guitar.NewThreadNum(i)
		if err != nil {
			panic(err)
		}
		threads[*threadNum] = guitar.NewThread(
			threadNum.RegularOpenNote().Up(turn[i]),
		)
	}

	return guitar.NewTurning(threads)
}

func printTurning(tn guitar.Turning) {
	threads := tn.Threads()
	nums := tn.ThreadNums()
	for _, num := range nums {
		fmt.Printf(
			"num: %d, note: %s\n",
			num,
			threads[num].OpenNote(),
		)
	}
}
