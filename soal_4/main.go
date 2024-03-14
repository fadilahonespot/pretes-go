package main

import "fmt"

func main() {
	fmt.Println(rekursifFaktorial(5))
}

func rekursifFaktorial(angka int) int {
	if angka == 1 {
		return 1
	}
	return angka * rekursifFaktorial(angka-1)
}
