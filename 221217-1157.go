package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	str, err := r.ReadString('\n')

	if err!=nil{
		log.Fatal(err)
	}

	var slice []int = make([]int, 27)

	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)

	for _,s := range str{
		slice[int(s)-64] += 1
	}

	var isSame bool
	for i, n := range slice{
		if n > slice[slice[0]]{
			isSame = false
			slice[0] = i
		} else if n == slice[slice[0]]{
			isSame = true
			slice[0] = i
		}
	}

	if isSame {
		fmt.Println("?")
	} else{
		fmt.Printf("%c\n", slice[0]+64)
	}
}
