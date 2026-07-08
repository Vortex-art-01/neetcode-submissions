func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}
 
	m := map[byte]int{}

	for i:=0; i < len(s); i++ {
		m[s[i]]++ 
	}

	for i:=0; i < len(t); i++ {
		if _, found := m[t[i]]; found {
			m[t[i]]--
		}
	}

	// for i:= 0; i < len(m); i++ {
		// if cnt, found := m[i]; found {
			// if cnt != 0 {
	// 			return false
	// 		}
	// 	}
	// }

	for _, cnt := range m {
		// fmt.Println("char, cnt", string(char) , cnt)
		if cnt != 0 {
			return false
		}
	}

	return true
}
