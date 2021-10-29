package main

import "fmt"

func main() {
	a := 0
	b := 1
	sum := 0
	c := 4000000
	for b < c {
		other := a + b
		if b%2 == 0 {
			sum = sum + b
		}
		a = b
		b = other
	}
	fmt.Println(sum)
}
