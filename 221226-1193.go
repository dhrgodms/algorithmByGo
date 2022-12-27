package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	var N, n, i int

	stdin := bufio.NewReader(os.Stdin)
	_, err := fmt.Scan(&N)

	if err != nil{
		fmt.Println(err)
		stdin.ReadString('\n')
	} else{
		n = int(math.Ceil((math.Pow(float64(1+8*N), 0.5) - 1)/float64(2)))
		i = N-1-n*(n-1)/2
		if n%2 == 0{
			fmt.Printf("%d/%d\n", i+1, n-i)
		}else{
			fmt.Printf("%d/%d\n", n-i, i+1)
		}
	}
}