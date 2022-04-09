package main

import "fmt"

func main() {
	var twoDimArray [3][4]float64
	twoDimArray[1][2] = 45
	twoDimArray[0][1] = 100

	for i, content := range twoDimArray {

		for j, value := range content {
			fmt.Printf("i: %v j: %v  value: %v\n ",i,j,value)
		}
		fmt.Println()
	}

}