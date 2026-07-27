func isPalindrome(s string) bool {
	// fmt.Println(s)
	prerared_s := strings.ReplaceAll(strings.ToLower(s), " ", "")
	prerared_s = strings.ReplaceAll(prerared_s, "?", "")
	prerared_s = strings.ReplaceAll(prerared_s, ".", "")
	prerared_s = strings.ReplaceAll(prerared_s, "!", "")
	prerared_s = strings.ReplaceAll(prerared_s, "/", "")
	prerared_s = strings.ReplaceAll(prerared_s, ",", "")
	prerared_s = strings.ReplaceAll(prerared_s, "'", "")
	prerared_s = strings.ReplaceAll(prerared_s, ":", "")
	// fmt.Println("prerared_s", prerared_s)

	l, r := 0, len(prerared_s)-1

	for l <= r {
		if prerared_s[l] != prerared_s[r] {
		// fmt.Println("error")
		// fmt.Println("s[l]", string(s[l]))
		// fmt.Println("s[r]", string(s[r]))
			
			return false
		}
		l++
		r--
	}
	return true
}
