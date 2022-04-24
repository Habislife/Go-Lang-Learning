package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "My name is fahad and my pet name is kitty"

	fmt.Println(strings.Contains(str, "fahad"))

	fmt.Println(strings.Count(str, "is"))
	
	fmt.Println(strings.HasPrefix(str, "My"))
	
	fmt.Println(strings.HasSuffix(str, "kitty"))

	fmt.Println(strings.Index(str, "fahad"))
	
	newStr := strings.Join([]string{"bishal","devkota", "habis"}," ")
	fmt.Println(newStr)

	fmt.Println(strings.Repeat(newStr, 5))

	fmt.Println(strings.ReplaceAll(str,"is","was"))

	fmt.Println(strings.Replace(str,"is","was",1))

	myStr := strings.Split(str, " ")
	fmt.Println(myStr[1])
	fmt.Println(len(str))
	fmt.Println(len(myStr))

	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
 }