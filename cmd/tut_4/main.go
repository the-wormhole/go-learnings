package main

import "fmt"

func main() {
	var arr [3]int32 // By default all values are zero meaning an array like [0,0,0]
	fmt.Println(arr[0])

	var arr1 [3]int32 = [3]int32{1, 2, 3} //Predefining an array
	fmt.Println(arr1[0])
	fmt.Println(arr1[1:3])

	arr2 := [3]int32{1, 2, 3}
	arr3 := [...]int32{1, 2, 3}

	fmt.Println(arr2, arr3)

	//Slices in Go are wrappers around arrays that give you a bunch of functionalities and features making it convenient for you to use them

	var slice []int32 = []int32{4, 5, 6}
	slice1 := []int32{1, 2, 3}

	fmt.Println(slice, slice1)

	fmt.Printf("Sice of Slice - %v and capacity - %v", len(slice), cap(slice))
	slice = append(slice, 7)
	fmt.Printf("\nSice of Slice - %v and capacity - %v", len(slice), cap(slice)) //Capacity got dohbled due to not space in the existing slice

	slice = append(slice, slice1...) //Append arrays using spread operator

	fmt.Println("\n", slice)

	var slice3 []int32 = make([]int32, 3, 8) //Declaring a slice using make() function, aand has a syntax as (type,len,capacity)

	fmt.Println(slice3)
}
