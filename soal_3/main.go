package main

func main() {
	printTriangle(5)
}

func printTriangle(n int) {
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			print("*")
		}
		println()
	}
}