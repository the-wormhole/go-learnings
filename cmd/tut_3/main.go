package main

import (
	"errors"
	"fmt"
)

func main() {
	var strPrint string = "You printed me!!"
	printString(strPrint)

	var temp1 int = 8
	var temp2 int = 6
	var resultDivision int = intDiv(temp1, temp2)

	fmt.Println(resultDivision)
	var resultRem int
	resultDivision, resultRem = retDivAndRem(temp1, temp2)

	fmt.Println(resultDivision, resultRem)

	fmt.Printf("Division result %v and remander is %v", resultDivision, resultRem)

	div, rem, err := retDivAndRemAndError(temp1, temp2) //Directly declared and assigend
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if rem == 0 {
	// 	fmt.Printf("\nDivision result %v", div)
	// } else {
	// 	fmt.Printf("\nDivision result %v and remander is %v", div, rem)
	// }

	switch { // Normal Switch statements in go
	case err != nil:
		fmt.Printf(err.Error())
	case rem == 0:
		fmt.Printf("\nDivision result %v", div)
	default:
		fmt.Printf("\nDivision result %v and remander is %v", div, rem)
	}
	switch rem { // Conditional Switch statements in go, like in other languages
	case 0:
		fmt.Printf("\nAccurate Division")
	case 1, 2:
		fmt.Printf("\nDivision was close")
	default:
		fmt.Printf("\nDivision was not close")
	}
}

func printString(str string) {
	fmt.Println(str)
}

func intDiv(num int, den int) int {
	var res int = num / den
	fmt.Println(res)

	return res
}

func retDivAndRem(num int, den int) (int, int) { // In go we can return multiple values at the same time
	return num / den, num % den
}

func retDivAndRemAndError(num int, den int) (int, int, error) {
	var err error //Default value is 'nil'
	if den == 0 {
		err = errors.New("Cannot Divide by Zero")

		return 0, 0, err
	}
	return num / den, num % den, err
}
