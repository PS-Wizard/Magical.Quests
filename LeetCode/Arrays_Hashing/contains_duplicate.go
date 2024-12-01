package arrayshashing

func containsDuplicate(nums []int) bool {
	got := make(map[int]struct{})
	for _, v := range nums {
		_, exists := got[v]
		if exists {
			return true
		}
		got[v] = struct{}{}
	}
	return false
}
