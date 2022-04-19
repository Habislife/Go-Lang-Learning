package main

import "fmt"

func main() {
	var summ float64
	summ = summation(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(summ)
	summ = summation(1, 2, 3, 4, 5)
	fmt.Println(summ)

}
func summation(inputParam ...float64) (sum float64) {
	sum = 0
	for _, entry := range inputParam {
		sum += entry
	}
	return sum
}