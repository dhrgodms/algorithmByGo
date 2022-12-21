package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var a, b, c int

	stdin := bufio.NewReader(os.Stdin)
	_, err := fmt.Scan(&a, &b, &c)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else{
		if b-c != 0{
		bep := (-1)*a/(b-c)+1
		if bep>0{
			fmt.Println(bep)
		} else{
			fmt.Println(-1)
		}
		} else{
			fmt.Println(-1)
		}
	}
}
