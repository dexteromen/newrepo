package main

import (
	"fmt"
)

func main() {
	// var x,y int
	// fmt.Println("Enter a number.")
	// fmt.Scan(&x,&y)
	// fmt.Println("Here is x,y:", x, y)

	// if x%2==0 {
	// 	fmt.Println("Number is Even")
	// }else{
	// 	fmt.Println("Number is Odd")
	// }
	// if y%2==0 {
	// 	fmt.Println("Number is Even")
	// }else{
	// 	fmt.Println("Number is Odd")
	// }

	// str := "ollo"
	// fmt.Println("Palindrome")
	// j:= len(str)-1
	// isPalindrome := true
	// for i := 0; i < len(str)/2; i++ {
	// 	if str[i] != str[j] {
	// 		isPalindrome = false
	// 		break
	// 	}
	// 	j--
	// }
	// if isPalindrome {
	// 	fmt.Println("The string is a palindrome.")
	// } else {
	// 	fmt.Println("The string is not a palindrome.")
	// }

	// var num int
	// fmt.Println("Enter a number:")
	// fmt.Scan(&num)

	// isPrime := true
	// if num <= 1 {
	// 	isPrime = false
	// } else {
	// 	for i := 2; i*i <= num; i++ {
	// 		if num%i == 0 {
	// 			isPrime = false
	// 			break
	// 		}
	// 	}
	// }

	// if isPrime {
	// 	fmt.Println(num, "is a prime number.")
	// } else {
	// 	fmt.Println(num, "is not a prime number.")
	// }

	// var n int
	// fmt.Println("Enter the number of terms:")
	// fmt.Scan(&n)

	var n int
	fmt.Println("Enter a number n:")
	fmt.Scan(&n)
	a, b := 0, 1
	fmt.Println("Fibonacci Series:")
	for i := 0; i <= n; i++ {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
	fmt.Println()
}
