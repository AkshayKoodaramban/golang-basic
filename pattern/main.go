package main

import "fmt"

func main() {
	var limit int
	fmt.Printf("enter the limit: ")
	fmt.Scanf("%d",&limit)
	for i:=1;i<=limit;i++{
		for j:=1;j<=limit;j++{
			if j<=i{
				fmt.Printf("%d",j)
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}