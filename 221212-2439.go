package main

import "fmt"

func main(){
	var a int
	var result string
	k:=0

	fmt.Scanf("%d", &a)
	for k<a {
		for j:=k; j>0;j-- {
			result += " "
		}
		for i:=0; i<k; i++{
			result += "*"
		}
		fmt.Println(result)
		k++
	}
}