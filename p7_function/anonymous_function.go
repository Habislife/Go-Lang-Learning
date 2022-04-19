package main

import "fmt"

func main() {
	greet := func(userName string) (greetings string) {

		return "Hello user " + userName + ".  Welcome to this tutorial"
	}
	fmt.Println(greet("Habis"))
	fmt.Println(greet("Bishal"))

}
