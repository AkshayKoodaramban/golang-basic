package main

import "fmt"

func main() {
	num := 1
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if j <= i {
				fmt.Printf("%d ", num)
				num++
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	
}
