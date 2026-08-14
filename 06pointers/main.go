package main

import "fmt"

func main() {
	// pointer - is a variable that stores the memory address of another variable
	// var ptr *int
	// fmt.Println(ptr)

	mynum :=29
	var ptr = &mynum //refrence operator
	fmt.Println(ptr)
	fmt.Println(*ptr)




}
