package main

import "fmt"

func main() {
	var R float64
	var n float64
	var P int
	fmt.Println("enter Principle Amount")
	fmt.Scanf("%d",&P)
	fmt.Println("enter your interest")
	fmt.Scanf("%f",&R)
	fmt.Println("number of years")
	fmt.Scanf("%f",&n)
	SI:=(float64(P)*R*n)/100
	fmt.Println("your simple interest is: ",SI)
}