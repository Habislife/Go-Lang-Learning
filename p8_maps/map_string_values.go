package main

import "fmt"

func main() {
	fruitData := map[string]int{
		"apples":   22,
		"oranges": 34,
	}
	// fmt.Println(fruitData["apples"])
	fruitData["oranges"]= 64
	 fruitData["bananas"] = 55
	// fmt.Println(fruitData)

	for fruit,quantity := range fruitData{
		fmt.Println("Fruits: ",fruit)
		fmt.Println("Quantity: ", quantity)
	}
	fruitData["watermelon"]=555
	fmt.Println("Before",len(fruitData))
	delete(fruitData,"bananas")
	fmt.Println("After",len(fruitData))


}