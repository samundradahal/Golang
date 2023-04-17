//While running packages we should run both the files (In this case, "go run main.go greetings.go")
// Variables shoul be decleared outside the function i.e global to access on another go file

package main

import "fmt"

var score = 66.7

func main() {
	greetings("Samundra")

	for _, v := range name {
		fmt.Println(v)
	}
	scoreof()

}
