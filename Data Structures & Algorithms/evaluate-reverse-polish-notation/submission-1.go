func evalRPN(tokens []string) int {
	stack := []int{}
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	// fmt.Println("tokens", tokens);

	for i := 0; i < len(tokens); i++ {
		if _, found := operators[tokens[i]]; found {
			// fmt.Println("found", tokens[i])
			second_num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			first_num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var res int

			if tokens[i] == "+" {
				res = first_num + second_num
			} else if tokens[i] == "-" {
				res = first_num - second_num
			} else if tokens[i] == "*" {
				res = first_num * second_num
			} else if tokens[i] == "/" {
				res = first_num / second_num
			}

			// fmt.Println("res", res)

			stack = append(stack, res)

		} else {
			// fmt.Println("else", tokens[i])
			intval, err := strconv.Atoi(tokens[i])
			if (err != nil) {
				fmt.Println("err", err)
				fmt.Println("intval", intval)
			}

			stack = append(stack, intval)
		}
	}

	// fmt.Println("stack", stack)

	return stack[0]
}
