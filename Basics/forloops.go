package main

import "fmt"

func main() {
	// x := 0

	// 	for x < 5 { //While loop
	// 		fmt.Println("Value of x:", x)
	// 		x++
	// 	}

	// 	for i := 0; i<5 ; i++{   //Normal for loop
	// 		fmt.Println("Value of i:", i)
	// 	}

	//for loop in slices

	name := []string{"Samundra", "Sagar", "Samir", "Sachin"}

	for index, value := range name {
		fmt.Println("Index is ", index, "and name is ", value)
	}

	for _, value := range name { // Using _ if variable or value is not needed
		fmt.Println(" Name is ", value)
	}

}
