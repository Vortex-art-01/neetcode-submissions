func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlpha(rune(s[l])) {
			l++
		}
		for r > l && !isAlpha(rune(s[r])) {
			r--
		}
		if l < r && strings.ToLower(string(s[l]))  != strings.ToLower(string(s[r])) {
			// fmt.Println("err")
			// fmt.Println("rune(s[l])",string(s[l]))
			// fmt.Println("rune(s[r])",string(s[r]))
			return false
		}

		l++
		r--
	}
	return true
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
