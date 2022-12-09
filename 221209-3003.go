package main

import "fmt"

func main(){
	var k, q, l, b, n, p int

	fmt.Scanf("%d %d %d %d %d %d", &k, &q, &l, &b, &n, &p)
	fmt.Printf("%d %d %d %d %d %d\n", 1-k, 1-q, 2-l, 2-b, 2-n, 8-p)
}
