package main

import "fmt"

func main() {
	fmt.Println("Tutorials")
	var (
		num1 int32
		num2 int64
	)
	num1 = 36
	num2 = 72
	fmt.Println(int64(num1) + num2)
}