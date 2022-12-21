package algorithm

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	var num1, num2 string
	var num1Result, num2Result int

	stdin := bufio.NewReader(os.Stdin)
	_, err := fmt.Scan(&num1, &num2)

	if err != nil{
		fmt.Println(err)
		stdin.ReadString('\n')
	}else {
		for i := 0; i < 3; i++ {
			num1Result += (int(num1[i])-48) * int(math.Pow(10, float64(i)))
			num2Result += (int(num2[i])-48) * int(math.Pow(10, float64(i)))
		}

		if num1Result > num2Result {
			fmt.Println(num1Result)
		} else {
			fmt.Println(num2Result)
		}
	}
}
