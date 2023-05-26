package main

import "fmt"

func main() {
	var limit int
	fmt.Println("enter array limit: ")
	fmt.Scanln(&limit)

	var sum int
	for i := 1; i <= limit; i++ {
		if i%2 != 0 {
			sum = sum + i
		}

	}
	fmt.Println(sum)
}
