package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	p, q, t := 0, 0, 0
	for x != 0 {
		p = x % 10
		q = x / 10
		if q != 0 {
			t = t*10 + p*10
		} else {
			t = t + p
		}
		x = q
	}
	fmt.Println(t)
	return temp == t
}
