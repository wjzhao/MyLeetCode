package main

// [1, 2, 3, 4, 5]

func maxProfit(prices []int) int {
	minPrice, result := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > result {
			result = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return result
}
