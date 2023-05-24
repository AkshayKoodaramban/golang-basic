package main

import (
	"fmt"
	"sort"
)

func main() {
	var limit int
	fmt.Println("enter array limit")
	fmt.Scanln(&limit)

	array1 := make([]int, limit)

	fmt.Println("enter array element:")
	for i := 0; i < limit; i++ {
		fmt.Scanln(&array1[i])
	}

	count := 0
	for i := 0; i < limit; i++ {
		if array1[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("number of even number is:", count)
}
