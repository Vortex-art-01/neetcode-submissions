func largestRectangleArea(heights []int) int {
	maxRect := 0
	stack := stack{
		arr: [][]int{},
	}

	for idx, height := range heights {

		if stack.IsEmpty() || stack.TopHeight() <= height {
			stack.Push(idx, height)
		} else {
			var last []int

			for !stack.IsEmpty() && stack.TopHeight() > height {
				last = stack.Pop()
				width := idx - last[0]
				localRect := width * last[1]

				if maxRect < localRect {
					maxRect = localRect
				}
			}

			stack.Push(last[0], height)
		}
	}


	for !stack.IsEmpty() {
		last := stack.Pop()
		width := len(heights) - last[0]
		localRect := width * last[1]

		// fmt.Println("last", last)
		// fmt.Println("width", width)
		// fmt.Println("localRect", localRect)
		if maxRect < localRect {
			maxRect = localRect
		}
	}
	return maxRect
}

type stack struct {
	arr [][]int
}

func (s *stack) Push(idx, height int) {
	s.arr = append(s.arr, []int{idx, height})
}

func (s *stack) Pop() []int {
	res := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1] 
	return res
}

func (s stack) Top() []int {
	return s.arr[len(s.arr)-1]
}

func (s stack) TopHeight() int {
	return s.arr[len(s.arr)-1][1]
}
func (s stack) IsEmpty() bool {
	return len(s.arr) == 0
}
