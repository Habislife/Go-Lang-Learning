package main

import "fmt"

func main() {
	arr1 := [3]int{
		20,
		30,
		40,
	}
	fmt.Println("Orginal array: ",arr1)
	modifyArray(&arr1)
	fmt.Println("Modified array: ",arr1)
}

func modifyArray(arr *[3]int){
	(*arr)[1] = 100 //arr[0]
	arr[2] = 400


}