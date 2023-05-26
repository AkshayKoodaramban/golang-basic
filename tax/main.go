package main

import "fmt"

func main() {
	var num float64
	var tax float64
	fmt.Println("enter the anual income")
	fmt.Scanln(&num)
	if num<=250000{
		fmt.Println("you have no Tax")
	}else if num>250000 && num <=500000{
		tax=(num*5)/100
		fmt.Println("Income tax amount =",tax )
	}else if num>500000 && num <=1000000{
		tax=(num*20)/100
		fmt.Println("Income tax amount =",tax )
	}else if num>1000000 && num <=5000000{
		tax=(num*30)/100
		fmt.Println("Income tax amount =",tax )
	}
	
}