package main

import "fmt"

func main() {
	fmt.Println("enter your input")
	var b int
	fmt.Scanln(&b) //getting the user input
	sum := 0

	for a := 1; a <= b; a++ {
		if a%2 == 0 {
			sum = sum + a
		}

	}
	fmt.Println(sum)
}
