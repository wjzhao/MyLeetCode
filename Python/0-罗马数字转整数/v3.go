package main

import "fmt"

func romanToInt(s string) int {
	chars := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	ans := 0
	for i := range s {
		if i < len(s)-1 && chars[string(s[i])] < chars[string(s[i+1])] {
			ans -= chars[string(s[i])]
		} else {
			ans += chars[string(s[i])]
		}
	}
	return ans
}

func main() {
	ret1 := romanToInt("III")
	ret2 := romanToInt("IX")
	ret3 := romanToInt("LVIII")
	fmt.Printf("%d %d %d", ret1,ret2,ret3)
}
