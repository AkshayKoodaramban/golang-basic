package main

import (
	"fmt"
)

func main() {
	var mark float64
	fmt.Println("enter your total mark Out of 100:")
	fmt.Scanf("%f", &mark)
	switch {
	case mark >= 90 && mark <= 100:
		fmt.Println("A Grade")
	case mark >= 80 && mark <= 89:
		fmt.Println("B Grade")
	case mark >= 70 && mark <= 79:
		fmt.Println("C Grade")
	case mark >= 60 && mark <= 69:
		fmt.Println("D Grade")
	case mark >= 50 && mark <= 59:
		fmt.Println("E Grade")
	case mark < 50:
		fmt.Println("You Are Failed")
	default:
		fmt.Println("Invalid data")

	}
}
