package main

func reverseWord(b []byte) []byte {
	i, j := 0, len(b)-1
	for j > i {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return b
}

func findBroad(b []byte) (int, int) {
	l, r := 0, len(b)-1
	for b[l] == ' ' || b[r] == ' ' {
		if b[l] == ' ' {
			l += 1
		}
		if b[r] == ' ' {
			r -= 1
		}
	}
	return l, r
}

func removeBlank(b []byte) []byte {
	i := 0
	for j := 0; j < len(b); j++ {
		if b[j] != ' ' {
			b[i] = b[j]
			i++
		} else if b[j+1] == b[j] && b[j+1] == ' ' {
			continue
		} else {
			b[i] = ' '
			i++
		}
	}
	return b[:i]
}

func reverseWords(s string) string {
	ss := []byte(s)
	left, right := findBroad(ss)
	ss_reverse := reverseWord(ss[left : right+1])
	sss_reverse := removeBlank(ss_reverse)
	i := 0
	for j := 0; j < len(sss_reverse); j++ {
		if sss_reverse[j] == ' ' {
			reverseWord(sss_reverse[i:j])
			i = j + 1
		}
		if j == len(sss_reverse)-1 {
			reverseWord(sss_reverse[i : j+1])
		}
	}
	return string(sss_reverse)
}
