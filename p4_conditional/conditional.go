package main

import "fmt"

func main() {
	num1 := 150
	num2 := 150

	if num1 < num2 {
		fmt.Println(num1, " < ", num2)
		}else if(num1 == num2){
			fmt.Println(num2," == ", num1)
				}else{
fmt.Println(num2," > ", num1)
	}
}