package play

import (
	"fmt"
	"testing"

	"github.com/sawadyecma/irreguitar/chord"
	"github.com/sawadyecma/irreguitar/guitar"
	"github.com/sawadyecma/irreguitar/sound"
	"github.com/sawadyecma/irreguitar/turning"
)

var turnings = map[string]map[guitar.ThreadNum]int{
	"regular":   {6: +0, 5: +0, 4: +0, 3: +0, 2: +0, 1: +0},
	"half down": {6: -1, 5: -1, 4: -1, 3: -1, 2: -1, 1: -1},
	"half up":   {6: +1, 5: +1, 4: +1, 3: +1, 2: +1, 1: +1},
	"dropD":     {6: -2, 5: +0, 4: +0, 3: +0, 2: +0, 1: +0},
	"DADGAD":    {6: -2, 5: +0, 4: +0, 3: +0, 2: -2, 1: -2},
	"DAEAC#E":   {6: -2, 5: +0, 4: +2, 3: +2, 2: +2, 1: +0},
	"FACGCE":    {6: +1, 5: +0, 4: -2, 3: +0, 2: +1, 1: +0},
	"DADAC#E":   {6: -2, 5: +0, 4: +0, 3: +2, 2: +2, 1: +0},
}

type playChordTestArg struct {
	rootPress Press
	presses   []Press
}
type playChordTest struct {
	arg  playChordTestArg
	want string
}

func test_play_chord(t *testing.T, play Play, tests []playChordTest) {
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("test:%d", i+1),
			func(t *testing.T) {

				got := play.
					Chord(
						tt.arg.rootPress,
						tt.arg.presses,
					).
					Name()

				if got != tt.want {
					t.Errorf("error => got: %s, want: %s", got, tt.want)
				}
			},
		)
	}
}

func newTurning(turn map[guitar.ThreadNum]int) turning.Turning {
	threads := make(map[guitar.ThreadNum]guitar.Thread, turning.DefaultThreadCnt)

	for i := 1; i <= turning.DefaultThreadCnt; i++ {
		threadNum := guitar.NewThreadNum(i)
		threads[*threadNum] = guitar.NewThread(
			threadNum.RegularOpenNote().Up(turn[*threadNum]),
		)
	}

	return turning.NewTurning(threads)
}

func Test_play_Chord_regular(t *testing.T) {
	play := NewPlay(
		newTurning(turnings["regular"]),
		sound.NewChordGenerator(
			chord.NewChordParser(),
		),
	)

	tests := []playChordTest{
		{
			arg: playChordTestArg{
				rootPress: Press{threadNum: *guitar.NewThreadNum(5), flet: 3},
				presses: []Press{
					{threadNum: *guitar.NewThreadNum(4), flet: 4},
					{threadNum: *guitar.NewThreadNum(3), flet: 0},
					{threadNum: *guitar.NewThreadNum(2), flet: 1},
					{threadNum: *guitar.NewThreadNum(1), flet: 0},
				},
			},
			want: "C",
		},
		{
			arg: playChordTestArg{
				rootPress: Press{threadNum: *guitar.NewThreadNum(5), flet: 5},
				presses: []Press{
					{threadNum: *guitar.NewThreadNum(4), flet: 7},
					{threadNum: *guitar.NewThreadNum(3), flet: 5},
					{threadNum: *guitar.NewThreadNum(2), flet: 6},
					{threadNum: *guitar.NewThreadNum(1), flet: 5},
				},
			},
			want: "Dm7",
		},
	}
	test_play_chord(t, play, tests)
}

func Test_play_Chord_DAEADbE(t *testing.T) {
	play := NewPlay(
		newTurning(turnings["DAEAC#E"]),
		sound.NewChordGenerator(
			chord.NewChordParser(),
		),
	)

	tests := []playChordTest{
		{
			arg: playChordTestArg{
				rootPress: Press{threadNum: *guitar.NewThreadNum(5), flet: 5},
				presses: []Press{
					{threadNum: *guitar.NewThreadNum(4), flet: 5},
					{threadNum: *guitar.NewThreadNum(3), flet: 4},
					{threadNum: *guitar.NewThreadNum(2), flet: 5},
				},
			},
			want: "DM7",
		},
	}
	test_play_chord(t, play, tests)
}
