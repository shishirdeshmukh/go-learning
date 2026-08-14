package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Hello , Craig"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating")

	//comma ok || error err

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	fmt.Printf("Type of input is %T \n", input)

}
