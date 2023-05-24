package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		sum := i * 5
		fmt.Printf("%d*5=%d \n", i, sum)
	}
}
