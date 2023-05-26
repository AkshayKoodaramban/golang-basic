package main

import (
	"fmt"
	// "sort"
)

func main() {
	var limit int
	fmt.Println("enter array limit :")
	fmt.Scanln(&limit)

	array := make([]int, limit)
	fmt.Println("enter array element")
	for i := 0; i < limit; i++ {
		fmt.Println("element is: ")
		fmt.Scanln(&array[i])
	}

	fmt.Println("enterd array is: ", array)

	// sort.Ints(array)
	
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}

	}
	fmt.Println("decending order array is: ", array)

	// // descending order
	// for i,j:=0,limit-1;i<j;i,j=i+1,j-1{
	// 	array[i],array[j]=array[j],array[i]
	// }
	// fmt.Println("descending order of array: ", array)
	
}
