package main

/*
Go is statically typed and strongly typed. It means that the compiler knows what kind of data will be stored in each variable at compile time.
x := 42        // int
y := 3.14      // float64
z := "golang"  // string
*/

import "fmt"

const Token string = "wirdovalue" // package level variable
// release := "2025" //  not allowed

func main() {
	// var username string = "admin"
	// fmt.Println("Username:", username)
	// fmt.Printf("Type of %T \n", username)

	// var isLoggedInb bool = true
	// fmt.Println(isLoggedInb)
	// fmt.Printf("Type of %T \n", isLoggedInb)

	// var smallval uint8 = 255
	// fmt.Println(smallval)
	// fmt.Printf("Type of %T \n", smallval)

	// var smallfloat float64 = 255.342222222121
	// fmt.Println(smallfloat)
	// fmt.Printf("Type of %T \n", smallfloat)

	var x int = 10  // explicit type
	var y = "hello" // type inferred
	z := 3.14       // shorthand (inside functions only)
	fmt.Println(x, y, z)

	// default values and some aliases
	// var av int
	// 	var b string    // ""
	// var c bool      // false
	// var d float64   // 0.0
	// var e []int     // nil
	// fmt.Println(av) //0
	// fmt.Printf("Type of %T \n", av)

	//implicit type
	// var web = "goodboy.com"
	// fmt.Println(web)

	// no var style
	// noofuser := 32322
	// fmt.Println(noofuser)

	a, b := 1, 2
	a, d := 5, 6 // redeclaration of a is allowed only when there is a new variable
	fmt.Println(a, b, d)

	fmt.Println(Token)
}
