package main

import "fmt"

func main() {
	var limit int
	fmt.Println("enter the number of rows &column")
	fmt.Scanln(&limit)
	array1 := make([][]int, limit)
	for i := 0; i < limit; i++ {
		array1[i] = make([]int, limit)
	}

	fmt.Println("enter the first combonent of array")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanln(&array1[i][j])
		}
	}
	fmt.Println("array", array1)

}
