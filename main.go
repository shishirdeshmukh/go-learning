package main

import "fmt"

func main() {
	// userID := 10
	// fmt.Println("Hello, Go! , userID:", userID)
	// fmt.Println(add(5, 3))

	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	// fmt.Println("result:", result)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// func add(a int, b int) int {
// 	return a + b
// }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
