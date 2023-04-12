package mexicanWave_test

import (
	"reflect"
	"testing"

	"vanmoof/internal/mexicanWave"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"empty case":                        {input: "", want: []string{}},
		"proper wave no extras":             {input: "mary", want: []string{"Mary", "mAry", "maRy", "marY"}},
		"please everybody get seated first": {input: "mArY", want: []string{"Mary", "mAry", "maRy", "marY"}},
		"skip empty seats":                  {input: "mA  r.Y", want: []string{"Ma  r.y", "mA  r.y", "ma  R.y", "ma  r.Y"}},
		"let's include some numbers":        {input: "mA 345  r.Y1", want: []string{"Ma 345  r.y1", "mA 345  r.y1", "ma 345  R.y1", "ma 345  r.Y1"}},
		"definitely not my password":        {input: "p@55  w0rD:`)", want: []string{"P@55  w0rd:`)", "p@55  W0rd:`)", "p@55  w0Rd:`)", "p@55  w0rD:`)"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mexicanWave.Wave(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Wave(%s) = %v, want: %v", tc.input, got, tc.want)
			}
		})
	}
}
