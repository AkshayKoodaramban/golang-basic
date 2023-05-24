package main

import "fmt"

func main() {
	var limit int
	fmt.Println("enter array limit")
	fmt.Scanln(&limit)

	//make an array
	array1 := make([]int, limit)
	//add values to first array
	fmt.Println("enter first array element")
	for i := 0; i < limit; i++ {
		fmt.Scanln(&array1[i])
	}

	//make second array

	array2 := make([]int, limit)
	//add values to second array
	fmt.Println("enter second array element")
	for i:= 0; i < limit; i++ {
		fmt.Scanln(&array2[i])
	}
	fmt.Println("first array is :", array1)
	fmt.Println("second array is :", array2)

	fmt.Println("arrays after swapping")

	for i:=0;i<limit;i++{
		array1[i],array2[i]=array2[i],array1[i]
	}
	fmt.Println("first array is :", array1)
	fmt.Println("second array is :", array2)
}
