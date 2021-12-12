package chord

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sawadyecma/irreguitar/sound"
)

func Test_chordParser_Parse(t *testing.T) {
	p := chordParser{}

	tests := []struct {
		root  sound.Absnote
		notes []sound.Absnote
		want  string
	}{
		{sound.C3, []sound.Absnote{}, "C"},
		{sound.C3, []sound.Absnote{sound.E3, sound.G3}, "C"},
		{sound.C3, []sound.Absnote{sound.Eb3, sound.G3}, "Cm"},
		{sound.C3, []sound.Absnote{sound.E3, sound.G3, sound.B3}, "CM7"},
		{sound.C3, []sound.Absnote{sound.E3, sound.G3, sound.Bb3}, "C7"},
		{sound.C3, []sound.Absnote{sound.Eb3, sound.G3, sound.Bb3}, "Cm7"},
		{sound.C3, []sound.Absnote{sound.E2, sound.E3, sound.G3}, "Cadd9"},
	}

	for i, test := range tests {
		t.Run(
			fmt.Sprintf("test:%d", i+1),
			func(t *testing.T) {
				got := p.Parse(test.root, test.notes)
				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("want: %s, got: %s", test.want, got)
				}
			},
		)
	}
}
