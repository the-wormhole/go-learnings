package main

import "fmt"

func main() {
	var strPrint string = "You printed me!!"
	printString(strPrint)
}

func printString(str string) {
	fmt.Println(str)
}

func intDiv(num int, den int) int {
	var res int = num / den
	fmt.Println(res)

	return res
}
