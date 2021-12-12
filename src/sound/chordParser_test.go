package sound

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_chordParser_Parse(t *testing.T) {
	p := chordParser{}

	tests := []struct {
		root  Absnote
		notes []Absnote
		want  string
	}{
		{C3, []Absnote{E3, G3}, "C"},
		{C3, []Absnote{Eb3, G3}, "Cm"},
		{C3, []Absnote{E3, G3, B3}, "CM7"},
		{C3, []Absnote{E3, G3, Bb3}, "C7"},
		{C3, []Absnote{Eb3, G3, Bb3}, "Cm7"},
		{C3, []Absnote{E2, E3, G3}, "Cadd9"},
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

func Test_parsed_ChordType(t *testing.T) {
	tests := []struct {
		parsed parsed
		want   string
	}{
		{
			parsed: parsed{
				thirdType:   lenTypeMajor,
				seventhType: lenTypeMajor,
			},
			want: "M7",
		},
		{
			parsed: parsed{
				thirdType:   lenTypeMajor,
				seventhType: lenTypeMinor,
			},
			want: "7",
		},
		{
			parsed: parsed{
				thirdType:   lenTypeMajor,
				seventhType: lenTypeNone,
			},
			want: "",
		},
	}

	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("test:%d", i+1),
			func(t *testing.T) {
				got := tt.parsed.ChordType()
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("want: %s, got: %s", tt.want, got)
				}
			},
		)
	}
}
