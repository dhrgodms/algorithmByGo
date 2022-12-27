package algorithm

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getIndex(N int) int{
	n := int(math.Ceil((3+(math.Pow(float64(12*N-3), 0.5)))/6))
	return n
}

func main(){
	var n int

	stdin := bufio.NewReader(os.Stdin)
	_, err := fmt.Scan(&n)

	if err != nil{
		fmt.Println(err)
		stdin.ReadString('\n')
	} else{
		fmt.Println(getIndex(n))
		}
	}
