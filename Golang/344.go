package main

func reverseString(s []byte) string {
	l, r := 0, len(s)-1
	for r > l {
		s[l], s[r] = s[r], s[l]
		r--
		l++
	}
	return string(s)
}
