package main

import "fmt"

func main() {
	var flag int
	var strings string
	fmt.Println("enter a string")
	fmt.Scanln(&strings)


	for i,j := 0,len(strings)-1; i < (len(strings))/2; i,j=i+1,j-1 {
		if strings[i]!=strings[j]{
			flag=1
			break
		}
	}

	if flag==0{
		fmt.Println("its a palindrome")
	}else if flag==1{
		fmt.Println("its not a palindrome")
	}


}