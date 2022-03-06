package main

import "fmt"

func double(number *int) {
	*number *= 2
}

func createFloatPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	var myFloatPointer = createFloatPointer()
	fmt.Println(*myFloatPointer)

	var myBool bool = true
	printPointer(&myBool)

	amount := 6
	double(&amount)
	fmt.Println(amount)

	// var myInt int
	// var myIntPointer *int
	// myIntPointer = &myInt
	// fmt.Println(myIntPointer)

	// var myFloat float64
	// var myFloatPointer *float64
	// myFloatPointer = &myFloat
	// fmt.Println(myFloatPointer)

	// var myBool bool
	// myBoolPointer := &myBool
	// fmt.Println(myBoolPointer)

	// myInt := 4
	// myIntPointer := &myInt
	// fmt.Println(myIntPointer)
	// fmt.Println(*myIntPointer)

	// myFloat := 98.6
	// myFloatPointer := &myFloat
	// fmt.Println(myFloatPointer)
	// fmt.Println(*myFloatPointer)

	// myBool := true
	// myBoolPointer := &myBool
	// fmt.Println(myBoolPointer)
	// fmt.Println(*myBoolPointer)

	// myInt := 4
	// fmt.Println(myInt)
	// myIntPointer := &myInt
	// *myIntPointer = 8
	// fmt.Println(*myIntPointer)
	// fmt.Println(myInt)
}
