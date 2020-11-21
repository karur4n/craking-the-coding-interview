package q1_1

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	testcases := []struct{
		args  string
		want bool
	}{
		{"foo", false},
		{"asdf", true},
	}

	for _, tcase := range testcases {
		if got := IsUnique(tcase.args); got != tcase.want {
			t.Errorf("IsUnique(%s): %v, want %v", tcase.args, got, tcase.want)
		}
	}
}

