package main

import "fmt"

func main() {
	var limit int
	fmt.Println("enter array limit")
	fmt.Scanln(&limit)
	array := make([]int, limit)
	for i := 0; i < limit; i++ {
		fmt.Println("array element is")
		fmt.Scanln(&array[i])
	}
	result := make([]int, limit-1)

	for i := 0; i < (limit)-1; i++ {
		result[i] = array[i] * array[i+1]
	}
	fmt.Println("result array is")
	fmt.Println(result)

}
