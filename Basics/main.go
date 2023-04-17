package main

import "fmt"

func main() {

	//fmt.Println("Hello World")

	//strings

	var num1 string = "Hello"
	var num2 = "There"
	num3 := "Guys"

	fmt.Println(num1 + num2 + num3)

	// integers

	var num4 int = 34
	num5 := 22

	fmt.Println(num4, num5)

	// integer with define memory and bits

	var num6 int8 = 100
	//var num7 int8 = 128 // Error because the value ranges from -127 to +127
	fmt.Println(num6)

	//float

	var num7 = 67.784

	fmt.Println(num7)
}
