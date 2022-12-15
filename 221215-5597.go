package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var studentSlice []bool = make([]bool, 30)

	stdin := bufio.NewReader(os.Stdin)

	for i:=0;i<28;i++{
		var a int
		_, err := fmt.Scan(&a)

		if err != nil{
			fmt.Println(err)
			stdin.ReadString('\n')
		} else{
			studentSlice[a-1] = true
		}
	}

	for i:=0;i<30;i++{
		if studentSlice[i] == false{
			fmt.Println(i+1)
		}
	}
}
