package main

import (
	"fmt"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	frequency := [101]int32{}
	var max int32 = 0

	for i := 0; i < len(a); i++ {
		index := a[i]
		frequency[index]++
	}
	for i := 1; i <= 100; i++ {
		var result int32 = frequency[i] + frequency[i-1]
		if result > max {
			max = result
		}
	}
	return max;
}

func main() {
	// Write your code here
	var a = []int32{4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98, 1, 98, 98, 1, 1, 4, 98, 2, 98, 3, 9, 9, 3, 1, 4, 1, 98, 9, 9, 2, 9, 4, 2, 2, 9, 98, 4, 98, 1, 3, 4, 9, 1, 98, 98, 4, 2, 3, 98, 98, 1, 99, 9, 98, 98, 3, 98, 98, 4, 98, 2, 98, 4, 2, 1, 1, 9, 2, 4}
	result := pickingNumbers(a)
	fmt.Println(result)
}
