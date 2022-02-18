package main

func twoSum(array []int, target int) []int {
	m := make(map[int]int)
	for index, val := range array {
		another := target - val
		if _, ok := m[another]; ok {
			return []int{m[another], index}
		}
		m[val] = index
	}
	return nil
}
