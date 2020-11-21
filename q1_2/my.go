package q1_2

import (
	"sort"
	"strings"
)

func CheckPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aArray := strings.Split(a, "")
	bArray := strings.Split(b, "")

	sort.Strings(aArray)
	sort.Strings(bArray)

	for i := 0; i < len(a); i++ {
		if aArray[i] != bArray[i] {
			return false
		}
	}

	return true
}
