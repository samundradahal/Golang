package main

import "fmt"

func main() {

	//array(Having fixed size and same data type on all of the elements)

	var age [3]int = [3]int{30, 80, 54}
	var name = [4]string{"Samundra", "Sagar", "Kiran", "Samir"}

	fmt.Println(age, len(age), name, len(name))

	//slices(Dynamic array can have of any size)

	score := []int{12, 34, 56, 78}
	score = append(score, 23)
	score[4] = 21

	fmt.Println(score)

	//Slice Range

	fmt.Println(score[1:3])

}
