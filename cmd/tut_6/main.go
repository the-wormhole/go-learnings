package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8 //default values are set to zero, {mpg:0,gallons:0}
	gallons uint8
	// ownerInfo owner
	// owner // this also works
	//int //declares a field name int of type int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

//	type owner struct {
//		name string
//	}
type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 { //Struct method
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 { //Struct method
	return e.kwh * e.mpkwh
}
func canMakeit(e engine, miles uint8) { //Normal function	// Now we can pass any object to the first parameter given that the object has a milesLeft() method
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!!")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func main() {

	var newEngine gasEngine = gasEngine{8, 9}
	// owner{"Nayan"}} // or {mpg:8,gallons:9}
	newEngine.gallons = 20 //works
	fmt.Println(newEngine.gallons, newEngine.mpg)
	// newEngine.name)
	// newEngine.ownerInfo.name
	fmt.Printf("Total miles left in the tank - %v", newEngine.milesLeft())

	fmt.Println("\nCan I make it to the destintion:")
	canMakeit(newEngine, 100)

	var newEngine1 electricEngine = electricEngine{8, 10}
	fmt.Println("\nCan I make it to the destintion:")
	canMakeit(newEngine1, 100)
}
