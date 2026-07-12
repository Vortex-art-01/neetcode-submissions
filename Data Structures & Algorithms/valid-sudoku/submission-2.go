func isValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    squares := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        squares[i] = make(map[byte]bool)
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == byte('.') {
                continue
            }

            val := board[r][c]
            sqIdx := (r/3)*3 + c/3

            if rows[r][val] ||
               cols[c][val] ||
               squares[sqIdx][val] {
                // fmt.Println("rows[r][val]", rows[r][val],"cols[c][val]", cols[c][val],"squares[sqIdx][val]", squares[sqIdx][val])
                // fmt.Println("r", r, "val", val, string(val))
                // fmt.Println("squares", squares)
                return false
               }

            rows[r][val] = true
            cols[c][val] = true
            squares[sqIdx][val] = true
        }
    }

    return true
}

