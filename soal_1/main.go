package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("radar"))
}

func isPalindrome(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-1-i] {
			return false
		}
	}
	return true
}