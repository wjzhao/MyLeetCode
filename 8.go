package main

import "fmt"

func myAtoi(s string) int {
	var sign = false
	for _, c := range s {
		// fmt.Println(sign)
		// if c == ' ' {
		// 	continue
		// }
		if c == '+' {
			sign = true
		} else {
			fmt.Println(sign)
		}
	}
	// fmt.Println(sign)
	return 0
}
