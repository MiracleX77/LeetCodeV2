package main

import (
	"fmt"
)

func main() {
	nums1 := []int{
		7, 1, 5, 3, 6, 4,
	}
	fmt.Println(maxProfit(nums1))
}

func maxProfit(prices []int) int {
	max_profit := 0
	buy_price := prices[0]

	for _, price := range prices {
		if price < buy_price {
			buy_price = price
		} else {
			profit := price - buy_price
			if profit > max_profit {
				max_profit = profit
			}
		}
	}

	return max_profit
}
