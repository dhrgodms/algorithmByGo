package algorithm

import "fmt"

func max(n1 int, n2 int, n3 int) int{
	var maxNum int
	if n1>=n2 {
		if n1>=n3{
		maxNum = n1
		} else{
			maxNum =n3
		}
	} else if n2>=n3 {
		maxNum=n2
	} else {
		maxNum=n3
	}
	return maxNum
}

func main()  {
	var a,b,c int
	var result int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a==b&&b==c {
		result = 10000+ a*1000
	} else if a==b||b==c {
		result = 1000 + b*100
	} else if a==c{
		result = 1000 + a*100
	} else{
		result = max(a,b,c)*100
	}
	fmt.Println(result)
}
