package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the insertionSort1 function below.
func countingSort(arr []int32) []int32 {
	var frequency = make([]int32, 100, 100)
	var sorted []int32

	for i := 0; i < len(arr); i++ {
		index := arr[i]
		frequency[index]++
	}
	for i, freq := range frequency {
		for j := 0; j < int(freq); j++ {
			sorted = append(sorted, int32(i))
		}
	}
	return sorted
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 2048*2048)

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

	result := countingSort(arr)
	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)
		if i != len(result)-1 {
			fmt.Printf(" ")
		}
	}
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
