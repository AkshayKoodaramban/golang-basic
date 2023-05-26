package main

import "fmt"

func main() {
	
	array1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	array2 := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}


	fmt.Println("Array 1:")
	displayArray(array1)
	fmt.Println("Array 2:")
	displayArray(array2)

	
	sumArray := addArrays(array1, array2)

	
	fmt.Println("Sum Array:")
	displayArray(sumArray)
}


func addArrays(array1, array2 [][]int) [][]int {
	rows := len(array1)
	cols := len(array1[0])

	sumArray := make([][]int, rows)
	for i := 0; i < rows; i++ {
		sumArray[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			sumArray[i][j] = array1[i][j] + array2[i][j]
		}
	}
	return sumArray
}

func displayArray(array [][]int) {
	rows := len(array)
	cols := len(array[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}
}
