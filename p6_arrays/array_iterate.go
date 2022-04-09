package main

import "fmt"

func main() {
	var arr = [5]float64{11, 22, 33, 44, 55}
	for _, content := range arr {
		fmt.Println(content)
	}
	var str1 = [5]string{"My", "name", "is", "Bishal", "Devkota"}
	for _, value := range str1 {
		fmt.Println(value)
	}
}