package main

import "fmt"

func main() {
	a := 1 + 2i
	fmt.Println(a)
	b := complex(2,4)
	fmt.Println(b)
	c := a + b
	fmt.Println(c)
	d := a * b 
	fmt.Println(d)
}