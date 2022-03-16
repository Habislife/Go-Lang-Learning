package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	num1 = 12
	num2 = 35
	fmt.Printf("%v + %v = %v\n",num1,num2,num1+num2)
	fmt.Printf("%v - %v = %v\n",num2,num1,num2-num1)
	fmt.Printf("%v * %v = %v\n",num1,num2,num1*num2)
	fmt.Printf("%v / %v = %v\n",num2,num1,num2/num1)

}