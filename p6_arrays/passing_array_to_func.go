package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	fmt.Println("Original Array:",arr1)

	modifiedArray := modifyArray(arr1)
	fmt.Println("modified Array:", modifiedArray)
}

func modifyArray(newArray [3] int) (arr[3] int){
	arr= newArray
	arr[1] = 100
	return arr 
}