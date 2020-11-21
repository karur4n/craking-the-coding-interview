package q1_1

func IsUniqueInBook(str string) bool {
	if len(str) > 128 {
		return false
	}

	charSet := [128]bool{false}

	for i := 0; i < len(str); i++ {
		if charSet[str[i]] {
			return false
		}
		charSet[str[i]] = true
	}

	return true
}
