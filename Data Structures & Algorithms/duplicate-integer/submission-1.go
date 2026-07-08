func hasDuplicate(nums []int) bool {
    m := map[int]int{}

    for _, v := range nums {
        if _, found := m[v]; found {
            return true
        } else {
            m[v] = 1
        }
    }
    return false
}
