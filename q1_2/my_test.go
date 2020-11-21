package q1_2

import "testing"

func TestCheckPermutation(t *testing.T) {
	testcases := []struct{
		args []string
		want bool
	}{
		{[]string{"foo", "bar"}, false},
		{[]string{"test", "estt"}, true},
		{[]string{"tomato", "otamot"}, true},
	}

	for _, tcase := range testcases {
		if got := CheckPermutation(tcase.args[0], tcase.args[1]); tcase.want != got {
			t.Errorf("CheckPermutation(%s): %v, want %v", tcase.args, got, tcase.want)
		}
	}
}