package Alg

func reSequence(s1, s2 string) bool {
	m1 := make(map[string]int)
	for _, v := range s1 {
		m1[string(v)] += 1
	}
	for _, v := range s2 {
		m1[string(v)] -= 1
		if m1[string(v)] < 0 {
			return false
		}
	}
	return true
}

//alternative approach: iterate through s1, check the count of each char in s1 have same appearance counts in s2
//slower than first approach
/*
func reSequence(s1, s2 string) bool {
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
*/
