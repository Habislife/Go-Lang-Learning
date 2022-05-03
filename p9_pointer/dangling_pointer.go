package main

import "fmt"

func main() {
	var ptr *string
	if ptr == nil {
		fmt.Println("It is dangling pointer")
	}
	var x string 
	x = "Fahad"
	ptr = &x
	// fmt.Printf("type of ptr is : %T and the value is : %v\n",ptr,ptr)
	// fmt.Printf("type of x is : %T  and the value is : %v\n",x,x)
	x = "ali"
	*ptr = "sarwar"
	fmt.Printf(x)
}