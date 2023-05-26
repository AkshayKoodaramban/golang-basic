package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}
func subtraction(a, b int) int {
	return a - b
}
func multiplication(a, b int) int {
	return a * b
}
func division(a, b float64) float64 {
	if b != 0 {
		return float64(a / b)
	} else {
		return 0
	}
}
func main() {
	var num1 int
	var num2 int
	fmt.Println("1 for addition")
	fmt.Println("2 for subtraction")
	fmt.Println("3 for multiplication")
	fmt.Println("4 for division")
	var num int
	fmt.Println("enter your action")
	fmt.Scanln(&num)

	if num > 0 && num < 5 {
		fmt.Println("enter your first value")
		fmt.Scanln(&num1)
		fmt.Println("enter your second value")
		fmt.Scanln(&num2)
		switch num {
		case 1:
			result := addition(num1, num2)
			fmt.Println("result = ", result)
		case 2:
			result := subtraction(num1, num2)
			fmt.Println("result = ", result)
		case 3:
			result := multiplication(num1, num2)
			fmt.Println("result = ", result)
		case 4:
			result := division(float64(num1), float64(num2))
			fmt.Println("result = ", result)
		default:
			fmt.Println("Invalid data")
		}
	} else {
		fmt.Println("Invalid data")
	}

}
