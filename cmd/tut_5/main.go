package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "subs" //String

	var index = str[0]
	fmt.Printf("%v,%T\n", index, index)
	str1 := []string{"s", "u", "b", "s"} //Array of Strings

	for i := range str1 {
		fmt.Println(str1[i])
		fmt.Println(str[i])
	}

	var str2 = []rune("test Rune")
	fmt.Printf("\nmy Rune = %v, my String = %v", str2, str1)

	var catStr = ""

	for i, v := range str1 {
		catStr += v
		catStr += str1[i] //With each of these operations we are actually creating a new string that gets assigned with this value, because natively string are immutable in Go. This is very inefficient so we use String Builder library in Go
	}

	fmt.Printf("\nConcatenated string = %v", catStr)

	var builtStr strings.Builder

	for i := range str1 {
		builtStr.WriteString(str1[i])
	}
	catStr = builtStr.String()
	fmt.Printf("\n%v", catStr)
}
