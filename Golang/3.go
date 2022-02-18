package main

func longestSubString(s string) int {
	n := len(s)
	right, current := 0, 0
	m := make(map[byte]int, n)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		current = max(current, right-i)
	}
	return current
}
