package arrayshashing

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for idx, num := range nums {
		m[num] = idx
	}

	for idx, num := range nums {
		if idy, ok := m[target-num]; ok && idx != idy{
			return []int{idx, idy}
		}
	}
	return []int{}
}
