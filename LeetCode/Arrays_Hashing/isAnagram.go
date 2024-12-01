package arrayshashing

func isAnagram(s string, t string) bool {
	var countS, countT [26]int
	for _, val := range s {
		countS[val-'a']++
	}
	for _, val := range t {
		countT[val-'a']++
	}
	return countS == countT
}

