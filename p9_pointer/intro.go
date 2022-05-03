package main

import "fmt"

func main() {
	var firstName string
	var ptrFirstName *string

	firstName = "Bishal"
	ptrFirstName = &firstName
	fmt.Println(firstName)
	fmt.Println(ptrFirstName)
	fmt.Println(*ptrFirstName)
}