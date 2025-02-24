package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter number 1: ")
	fmt.Scan(&num1)

	fmt.Print("Enter number 2: ")
	fmt.Scan(&num2)

	fmt.Println("Number 1:", num1)
	fmt.Println("Number 2:", num2)
}