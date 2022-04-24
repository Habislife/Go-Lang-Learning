package main

import "fmt"

func main() {
	var myStr string
	myStr = "Señor"
	//printBytes(myStr)
	//fmt.Println(len(myStr))
	//printRune(myStr)
	loopRune(myStr)

}

func printRune(str string){
	r := []rune(str)
	for i:= 0; i<len(r); i++{
		fmt.Printf("%c %x ",r[i],r[i])
		fmt.Println()
	}
}

func loopRune(str string){
	r := []rune(str)
	for index, value:= range r{
		fmt.Printf("%v %c %x\n",index,value,value)
	}
}

func printBytes(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
		fmt.Printf("%x ", str[i]) //print hex value
	}
	fmt.Println()
}