func dailyTemperatures(temperatures []int) []int {
	stack := stack{
		arr: []int{},
	}

	res := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		if stack.Empty() {
			stack.Push(i)
			continue
		}

		for !stack.Empty() && temperatures[stack.Top()] < temperatures[i] {
			top_idx := stack.Pop()
			res[top_idx] = i - top_idx
		}

		stack.Push(i)
	}

	return res

}

type stack struct {
	arr []int
}

func (s *stack) Top() int {
	return s.arr[len(s.arr)-1]
}

func (s *stack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *stack) Pop() int {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}

func (s stack) Empty() bool {
	return len(s.arr) == 0
}

