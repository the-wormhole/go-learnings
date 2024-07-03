package main

import (
	"fmt"
)

func main() {

	//var p *int32
	var p *int32 = new(int32) //creates a memory addess and assigns it to this variable
	var i int32

	p = &i
	fmt.Printf("Value of pointer - %v and variable %v", p, i)

	*p = 4
	fmt.Printf("\nValue of pointer - %v and variable %v", *p, i)

	//where as

	var k int32 = 5
	k = i //Copies the value to a new memory location
	fmt.Printf("\nValue of pointer - %v and variables %v %v", *p, i, k)

	var slice1 []int32 = []int32{1, 2, 3}
	var slice2 []int32 = slice1
	slice2[2] = 5 // updates both the slices because assignment operator creates a soft copy in case of Slices
	fmt.Printf("\nSlice 1 = %v", slice1)
	fmt.Printf("\nSlice 2 = %v", slice2)

	var arr1 [4]int32 = [4]int32{1, 2, 3, 4}
	fmt.Printf("\nThe original array = %v", arr1)

	var res [4]int32 = square(&arr1)
	fmt.Printf("\nThe new array is = %v and the old array = %v", res, arr1) //As it can be seen, when we pass the address of the array to the function, both original and the new arrays are updated
}

func square(arr2 *[4]int32) [4]int32 {

	for i := range arr2 {
		arr2[i] = arr2[i] * arr2[i]
	}
	return *arr2
}
