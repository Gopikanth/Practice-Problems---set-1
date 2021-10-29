package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("enter the input")
	var a int
	fmt.Scanln(&a)
	prime(a)
}
func prime(a int) {

	//here we are using trivial division algorithm
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			fmt.Println(i)
			a = a / i
			i--
		}
	}
	if a > 0 {
		fmt.Println(a)
	}

}
