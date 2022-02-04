package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	longest := ""
loop:
	for idx := range s[:len(s)-1] {
		if len(s)-idx < len(longest) {
			break loop
		}
		// for x := idx + 1; x < len(s)+1; x++ {
		for x := len(s); x-idx > len(longest); x-- {
			if x-idx < len(longest) {
				break
			}
			p := s[idx:x]
			// fmt.Println("Checking:", p)
			if isPalindrome(p) && len(p) > len(longest) {
				longest = p
				// fmt.Println("New longest:", longest)
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	s2 := ""

	for x := len(s) - 1; x >= 0; x-- {
		s2 += string(s[x])
	}

	return s == s2
}
