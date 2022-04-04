package main

import "fmt"

func main() {
	num1 := 50
	num2 := 100
	switch {
	case num1 == 50 && num2 < 100:
		fmt.Println("Num1 is 50")
	case num1 <50:
		fmt.Println("Num1 is less than 50")
	case num1 == 50 && num2  >=50:
		fmt.Println("Num1 is 50 and Number2 is more than 50")
	default:
		fmt.Println("This is default case")
	} 
}