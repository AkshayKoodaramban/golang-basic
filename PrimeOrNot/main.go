package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter a number")
	fmt.Scanln(&num)

	if num > 1 {
		flag := 0
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = 1
				break

			}
		}
		if flag == 0 {
			fmt.Println("entered number is a Prime Number")
		} else {
			fmt.Println("entered number is not a PrimeNumber")
		}

	}else{
		fmt.Println("entered number is not a PrimeNumber")
	}
}
