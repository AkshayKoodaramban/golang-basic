package main

import (
	"fmt"

	
)

func main() {
	slice := make([]int, 10)
	slice[0] = 1
	slice[5] = 10
	slice[4] = 7
	fmt.Println(slice)

}