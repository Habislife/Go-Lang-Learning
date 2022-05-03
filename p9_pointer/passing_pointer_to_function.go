package main

import "fmt"

func main() {
	num1 := 100
	fmt.Println("Before passing to function: ", num1)
	//modifyVariable(&num1)
	simpleFunc(num1)
	fmt.Println("After passing to function: ", num1)

}

func modifyVariable(ptr *int) {
	*ptr = 200
}

func simpleFunc(num int) {
	num = 200

}
