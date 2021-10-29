package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter the number")
	var number int
	fmt.Scanln(&number)
	prime(number)
}

func prime(number int) []int{
	factors := make([]int ,0)
    for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
        if number%i == 0 {
            factors = append(factors, i)
            number/= i
            i--
        }
    }
    if number > 0 {
        fmt.Println(number)
    }
	return factors
}
func largest(factors []int) int {
    max := factors[0]
    for _, factor := range factors {
        if factor > max {
            max = factor
        }
    }
    return max
}