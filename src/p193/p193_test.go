package p193

import (
	"testing"
	. "testing/utils"
)

func TestValidPhoneNumbers(t *testing.T) {
	testCase := func(in string, want []string) {
		got := validPhoneNumbers(in)
		if !IsStrsEqual(got, want) {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase("file.txt", []string{
		"987-123-4567",
		"(123) 456-7890",
	})
}
