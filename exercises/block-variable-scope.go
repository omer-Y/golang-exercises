package main

import "fmt"

var packageVar = "package"

func main() {
	var functionVar = "function"

	if true {
		var conditionalVar = "conditional"
		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionalVar)
	}

	fmt.Println(packageVar)
	fmt.Println(functionVar)
}
