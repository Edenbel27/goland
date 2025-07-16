// Task 2
package main

import (
	"fmt"
)
func charCounter(){
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	charCount := make(map[rune]int)
	for _, char := range input{
		charCount[char]++
	}
	for char , count := range charCount{
		fmt.Printf("%c : %d\n", char, count)
	}
}

func palindromeCheck(){
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)
	i , j := 0 , len(input) - 1
	for i < j{
		if input[i] != input[j]{
			fmt.Println("The string is not a palindrome")
			return
		}
		i++
		j--
	}
	fmt.Println("The string is a palindrome")
}
func main() {
	charCounter()
	palindromeCheck()
}