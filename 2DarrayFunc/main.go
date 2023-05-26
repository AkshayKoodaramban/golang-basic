package main

import "fmt"

func main() {
	fmt.Println("Enter the size of the arrays:")
	var rows, cols int
	fmt.Print("Rows: ")
	fmt.Scanln(&rows)
	fmt.Print("Columns: ")
	fmt.Scanln(&cols)

	fmt.Println("Enter values for the first array:")
	array1 := getArray(rows, cols)

	fmt.Println("Enter values for the second array:")
	array2 := getArray(rows, cols)

	sumArray := addArray(array1, array2)

	fmt.Println("Sum of the arrays:")
	displayArray(sumArray)
}

func getArray(rows, cols int) [][]int {
	array := make([][]int, rows)
	for i := 0; i < rows; i++ {
		array[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Printf("Enter element [%d][%d]: ", i, j)
			fmt.Scanln(&array[i][j])
		}
	}
	return array
}

func addArray(array1, array2 [][]int) [][]int {
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
