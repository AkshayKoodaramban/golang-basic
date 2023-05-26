package main

import "fmt"

func main() {
	var wt float64
	var le float64
	var asgn float64
	fmt.Println("enter your written test: ")
	fmt.Scanln(&wt)
	fmt.Println("enter your lab exam mark: ")
	fmt.Scanln(&le)
	fmt.Println("enter your assignments mark: ")
	fmt.Scanln(&asgn)
	overall:=(wt*70)/100+(le*20)/100 + (92*10)/100
	fmt.Println("Grade of the student is : ",overall)
}