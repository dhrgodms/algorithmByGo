package algorithm

import "fmt"

func main(){
	var a int
	k:=1

	fmt.Scanf("%d", &a)

	for k<=a {
		result:=""
		for i:=a-k; i>0; i--{
			result += " "
		}
		for j:=1; j<=k; j++{
			result += "*"
		}
		fmt.Println(result)
		k++
	}
}