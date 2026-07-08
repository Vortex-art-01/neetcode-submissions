func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sorted_str := string(runes)
	// fmt.Println("str",str,"sorted_str", sorted_str)

		m[sorted_str] = append(m[sorted_str], str)
	}

	// fmt.Println(m)

	res := [][]string{}

	for _, val := range m {
		res = append(res, val)
	}
	return res
}
