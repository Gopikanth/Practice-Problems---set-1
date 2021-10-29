package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("enter")
	var num int
	fmt.Scanln(&num)
	prime(num)
}
func prime(num int){
if num<2{
	fmt.Println("Number should greater than 2 ")
	return
}
square:= int(math.Sqrt(float64(num)))
 for i:=2; i<=square; i++{
      if num % i == 0 {
         fmt.Println("Non Prime Number")
         return
      }
   }
   fmt.Println("Prime Number")
   return
}