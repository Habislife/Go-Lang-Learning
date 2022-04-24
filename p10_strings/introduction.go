package main

import "fmt"

func main() {
	var myStr string

	myStr = "Bishal"
	printBytes(myStr)
}

func printBytes(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c",str[i])
		fmt.Printf("%x",str[i]) //print hex value
	}
	fmt.Println()
}