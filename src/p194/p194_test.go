package p194

import (
	"testing"
	. "testing/utils"
)

func TestTransposeFile(t *testing.T) {
	testCase := func(in string, want []string) {
		got := transposeFile(in)
		if !IsStrsEqual(got, want) {
			t.Errorf("Case %v: expected %#v, got %#v", in, want, got)
		}
	}

	testCase("file.txt", []string{
		"name alice ryan",
		"age 21 30",
	})
}
