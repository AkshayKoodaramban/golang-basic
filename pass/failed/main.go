package main

import "fmt"

func main() {
	var mark float64
	fmt.Println("enter your name out of 100")
	fmt.Scanf("%f",&mark)
	if mark>100{
		fmt.Println("invalid data")
	}else if mark<50{
		fmt.Println("you are failed")
	}else{
		fmt.Println("you are passed")
	}
}