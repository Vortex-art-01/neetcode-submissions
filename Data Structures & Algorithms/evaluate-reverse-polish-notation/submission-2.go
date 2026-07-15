func evalRPN(tokens []string) int {
    index := len(tokens) - 1

    var dfs func() int
    dfs = func() int {
        token := tokens[index]
        index--

        if token != "+" && token != "-" && token != "*" && token != "/" {
            val, _ := strconv.Atoi(token)
            return val
        }

        right := dfs()
        left := dfs()

        switch token {
        case "+":
            return left + right
        case "-":
            return left - right
        case "*":
            return left * right
        default:
            return left / right
        }
    }

    return dfs()
}