package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// var temp float32 = 3276758.6
	// var temp1 int = 255
	//int32
	//int64
	//uint16
	//uint32
	//uint64
	// fmt.Println(temp1)
	// var temp2 float32 = 140

	//fmt.Println(temp1 + temp2) // Not possible due to different types.

	var str string = "Hey"
	fmt.Println(str)

	var strConc string = "Hey" + " " + "Dude"
	fmt.Println(strConc)

	fmt.Println(len("test")) // Returns the number of bytes, since special characters store more than a single byte, it is not recommended to use this function

	fmt.Println(utf8.RuneCountInString("")) // For normal string length to be returned

	var testRune rune = 'a'
	fmt.Println(testRune)

	var myBool bool = true
	fmt.Println(myBool)

	// Default value for string is '', for boolean false and faor all int and floats 0

	temp := 37
	temp1, temp2 := 21, 13

	fmt.Println(temp, temp1, temp2)

	const myConst string = "well"
	fmt.Println(myConst)

}
