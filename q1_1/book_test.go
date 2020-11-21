package q1_1

import (
	"testing"
)

func TestIsUniqueInBook(t *testing.T) {
	testcases := []struct{
		args  string
		want bool
	}{
		{"foo", false},
		{"asdf", true},
	}

	for _, tcase := range testcases {
		if got := IsUniqueInBook(tcase.args); got != tcase.want {
			t.Errorf("IsUniqueInBook(%s): %v, want %v", tcase.args, got, tcase.want)
		}
	}
}

