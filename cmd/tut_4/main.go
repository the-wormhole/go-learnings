package main

import "fmt"

func main() {
	var arr [3]int32 // By default all values are zero meaning an array like [0,0,0]
	fmt.Println(arr[0])

	var arr1 [3]int32 = [3]int32{1, 2, 3} //Predefining an array
	fmt.Println(arr1[0])
	fmt.Println(arr1[1:3]) // Prints the 1 and 2 index values stored in the array

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

	var slice3 []int32 = make([]int32, 3, 8) //Declaring a slice using make() function, and has a syntax as (type,len,capacity). If no capacity value is passed, then len === capacity

	fmt.Println(slice3)

	// Maps in Go
	var map1 map[string]uint8 = make(map[string]uint8) //declares a map with of keys string type and unsigned integer as value
	fmt.Println(map1)

	map2 := map[string]uint8{"Adam": 23, "Sarah": 15} //Declaring with default value
	fmt.Println(map2["Adam"])
	fmt.Println(map2["Jacky"])

	map2["Bob"] = 14
	delete(map2, "Bob")

	var age, ok = map2["Adam"] //By default a map returns 2 values, 2nd being the true or false if the value is present in the map

	if ok {
		fmt.Printf("The age is:%v\n", age)
	} else {
		fmt.Println("Name not found!!")
	}

	//Loops

	for name := range map2 {
		fmt.Printf("Name:%v\n", name)
	}

	for name, age := range map2 {
		fmt.Printf("Name:%v and Age:%v\n", name, age)
	}

	for i, v := range arr2 {
		fmt.Printf("Index:%v,Value:%v\n", i, v)
	}

	// var i int32;
	for i := 0; i < len(arr2); i++ {
		fmt.Println(i, arr2[i])
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}
