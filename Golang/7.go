package main

func reverse7(x int) int {
	temp := 0
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
	}
	if temp > 1<<31-1 || temp < -(1<<31) {
		return 0
	}
	return temp
}
