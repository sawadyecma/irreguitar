package chord

import (
	"fmt"
	"reflect"
	"testing"
)

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
