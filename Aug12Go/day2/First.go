package main

import "fmt"

func calTotal(prices []int) int {
	total := 0
	for _, price := range prices {
		total = total + price
	}
	return total
}

func discountPrize(role string) int {
	switch role {
	case "admin":
		return 20
	case "user":
		return 10
	default:
		return 0
	}
}

func main() {
	prices := []int{1, 2, 3, 4, 5}
	total := calTotal(prices)
	discount := discountPrize("admin")
	fmt.Println("Total", total)
	fmt.Println("Discount", discount)
}
