package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Enter number 1: ")
	fmt.Scan(&x)

	fmt.Print("Enter number 2: ")
	fmt.Scan(&y)

	fmt.Println("Number 1:", x)
	fmt.Println("Number 2:", y)
}
