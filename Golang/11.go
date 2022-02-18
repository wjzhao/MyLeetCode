package main

func maxArea(height []int) int {
	n := len(height)
	area, w, h := 0, 0, 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			w = j - i
			if height[j] > height[i] {
				h = height[i]
			} else {
				h = height[j]
			}
			area = max(area, w*h)
		}
	}
	return area
}
