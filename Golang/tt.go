package main

import "fmt"

type Coin struct {
	val int
}

func coinBase(num int) int {
	// 0-正  1-反
	l := make([]Coin, num)
	for i := 0; i < len(l); i++ {
		l[i] = Coin{val: i + 1}
	}
	for i := 0; len(l) > 1; i++ {
		head := l[0]
		l = l[1:]
		if (i+1)%2 == 0 {
			l = append(l, head)
		}
		fmt.Println(l)
	}
	return l[0].val
}
