package algorithm

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var total, n int
	//fmt.Scanf("%d\n%d", &total, &n) 성공
	_, err := fmt.Scan(&total, &n) //성공
	stdin := bufio.NewReader(os.Stdin)

	if err!=nil{
		fmt.Println(err)
		stdin.ReadString('\n')
	} else{
		var result int
		for i:=0;i<n;i++ {
			var a, b int
			_, err := fmt.Scan(&a, &b)
			result += a * b
			if err != nil {
				fmt.Println(err)
				stdin.ReadString('\n')
				return
			}
		}
		if total == result {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
