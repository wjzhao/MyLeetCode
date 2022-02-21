package main

// "  abc def   hji "

func lengthOfLastWord(s string) int {
	l, r := 0, len(s)-1
	for s[r] == ' ' {
		r--
	}
	l = r
	for l >= 0 && s[l] != ' ' {
		l--
	}
	return r - l
}
