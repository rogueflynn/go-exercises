package main

/*
   in-place quicksort algorithm that uses the last index as the pivot
   and go lang swap syntax
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int32, low int, high int) int {
	pivot := arr[high]
	var i int = low - 1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			//swap
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)
}

// Complete the quickSort function below.
func quickSort(arr []int32, low int, high int) {
	if low < high {

		//Select pivot position
		pi := partition(arr, low, high)

		//sort elements on the left of the pivot
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	quickSort(arr, 0, int(n)-1)

	for _, resultItem := range arr {
		fmt.Printf("%d ", resultItem)
	}

	fmt.Printf("\n")
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
