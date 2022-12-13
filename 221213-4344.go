package main

import (
	"fmt"
	"math"
)

func main(){
	var c, k int
	fmt.Scan(&c)

	var ratioSlice []float64 = make([]float64, c)

	for k<c{
		var n, result, a int
		fmt.Scan(&n)

		var scoreSlice []int = make([]int, n)

		for i:=0; i<n; i++{
			fmt.Scan(&scoreSlice[i])
			result += scoreSlice[i]
		}
		average := int(result)/n

		for i:=0; i<n; i++{
			if scoreSlice[i]>average {
				a += 1
			}
		}
		tempAverage := math.Round(float64(a)/float64(n)*100000)*0.001
		ratioSlice[k] = tempAverage

		k++
	}

	for i:=0; i<c; i++{
		fmt.Printf("%.3f%%\n",ratioSlice[i])
	}
}
