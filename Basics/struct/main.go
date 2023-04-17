package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(info string, r *bufio.Reader) (string, error) {
	fmt.Print(info)
	input, err := r.ReadString('\n')
	return (strings.TrimSpace(input)), err

}
func CreateBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the bill name:")
	// name, _ := reader.ReadString('\n') //Input by the user
	// name = strings.TrimSpace(name)
	name, _ := input("Enter the bill name:", reader)
	b := NewBill(name)

	fmt.Println("Created the Bill:", name)
	return b

}

func addoptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := input("Choose an option : a- Add Item , s - Save Bill , t - Add Tip:", reader)
	switch opt { //Switch statement
	case "a":
		item, _ := input("Enter an item:", reader)
		price, _ := input("Enter a price:", reader)

		p, _ := strconv.ParseFloat(price, 64)
		b.additem(item, p)

		fmt.Println("Item Added:", item, price)
		addoptions(b)
	case "t":
		tip, _ := input("Enter a tip:", reader)

		t, _ := strconv.ParseFloat(tip, 64)
		b.addtip(t)
		fmt.Println("Tip Added:", tip)
		addoptions(b)
	case "s":
		b.save()
	default:
		fmt.Println("Please enter valid option")
		addoptions(b)
	}

}

func main() {
	mybill := CreateBill()
	fmt.Println(mybill)
	addoptions(mybill)
}
