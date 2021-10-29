package main

import "fmt"

func main() {
	a:= 1
	fmt.Println(a)
	b := 2
	fmt.Println(b)
	for i:=0 ; i<10 ; i++{
        c := a+b
		fmt.Println(c)
		a = b
		b = c
	} 
	
}