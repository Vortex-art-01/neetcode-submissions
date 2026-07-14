func isValid(s string) bool {
    stack := []byte{}
	hash := map[byte]byte {
		']': '[',
		'}': '{',
		')': '(',
	}
	
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 {
			lastOpen := stack[len(stack)-1]
				// fmt.Println("else 1", stack)
			stack = stack[:len(stack)-1]
				// fmt.Println("else 2", stack)

			char, found := hash[s[i]];
			if !found {
				// fmt.Println("false 1")
				return false
			}

			if char != lastOpen {
				// fmt.Println("false 2")
				return false
			}
		} else {
				// fmt.Println("false 3")
			return false
		}
	}

	if len(stack) > 0 {
						// fmt.Println("false 4")
		return false
	}

	return true
}
