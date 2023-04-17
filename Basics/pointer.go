package main

import "fmt"

func update(x *string) {
	*x = "Sagar"
}

func main() {
	name := "Samundra"

	fmt.Println("Memory Address of name is ", &name)
	ptr := &name // Memory Address of the variable
	fmt.Println(name)
	update(ptr)
	fmt.Println(name)

}
