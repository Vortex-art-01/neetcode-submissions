func isValidSudoku(board [][]byte) bool {

	for x := 0; x < 9; x++ {
		dlX := map[byte]int{}
		dlY := map[byte]int{}
	
		for y := 0; y < 9; y++ {
			if board[x][y] != '.' {
				if _, found := dlX[board[x][y]]; found {
		// fmt.Println("test", "val", string(board[x][y]), board[x][y], "y", y, "x", x, "dlX", dlX)
					return false
				}
				dlX[board[x][y]]++
			}
			if board[y][x] != '.' {
				if _, found := dlY[board[y][x]]; found {
		// fmt.Println("test 2")
					return false
				}
				dlY[board[y][x]]++
			}
		}
	
	}

	for y := 0; y <= 6; y+=3 {
		for x := 0; x <= 6; x+=3 {
			
			if (!check3x3(board, y, x)) {
			// if (!check3x3(board, 0, 0)) {
	// fmt.Println("test 3")
				return false
			}
		}
	} 

	return true
}

func check3x3(board [][]byte, xPlus, yPlus int) bool {

	dl := map[byte]int{}
	for y := 0 + yPlus; y < 3 + yPlus; y++ {

		for x := 0 + xPlus; x < 3 + xPlus; x++ {
			
			if board[y][x] != '.' {
				if _, found := dl[board[y][x]]; found {
					return false
				}
				dl[board[y][x]]++
			}

		}

	} 

	// fmt.Println("dl", dl)
	return true
}


