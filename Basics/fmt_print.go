package main

import "fmt"

func main() {
	//Print (Print on same line )

	fmt.Print("Hello,")
	fmt.Print("World\n")

	// Println (Automatically goes on newline)
	fmt.Println("Hello")
	fmt.Println("World")

	//Printf(Formatted Strings)

	var name = "Samundra Dahal"
	var age = 24

	fmt.Printf("My name is %v and my age is %v \n ", name, age)

	//Sprintf(Save formatted Strings)

	var str = fmt.Sprintf("My name is %v and my age is %v", name, age)
	fmt.Println("Saved string:", str)
}
