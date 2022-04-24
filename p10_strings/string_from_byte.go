package main

import "fmt"

func main() {
	bytes := []byte{0x48, 0x49}
	myStr := string(bytes)
	fmt.Println(myStr)

	decbytes := []byte{67, 97, 102, 195, 169}

	str := string(decbytes)
	fmt.Println(str)
	fmt.Println(len(str))
}