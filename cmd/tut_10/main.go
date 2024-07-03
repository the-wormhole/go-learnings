package main

import (
	"fmt"
)

func main() {

	var t1 float32 = 4
	var t2 float32 = 2
	var div = genericsTest(t1, t2)

	fmt.Printf("Division result = %v, Type of div = %T", div, div)
}

func genericsTest[T int | int32 | float32 | int64](num1 T, num2 T) T {

	var res = num1 / num2
	return res
}
