package main

/*
	Problem Solving Exercise Using go
	Calculates hourglass values of 6x6 matrix

	Input: text file in the format:

	    -9 -9 -9 1 1 1
	    0 -9 0 4 3 2
	    -9 -9 -9 1 2 3
	    0 0 8 6 6 0
	    0 0 0 -2 0 0
	    0 0 1 2 4 0

	    Read  from stdin

	Output: Absolute path of the file the result will be written to
		Stored in an environment variable call OUTPUT_PATH

	Usage:
		cat <INPUT_FILE> | go run hourglass.go
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func hourglassSum(arr [][]int32) int32 {
	var i, j int
	length := 6
	var max int32 = 0
	var init = true
	for i = 0; i < length; i++ {
		for j = 0; j < length; j++ {
			rowSlice := j + 3
			if rowSlice <= length && (i+2) < length {
				var sum int32 = 0
				hourGlassArray := make([]int32, 7)
				rowOne := make([]int32, length)
				rowTwo := make([]int32, length)
				rowThree := make([]int32, length)
				copy(rowOne, arr[i])
				copy(rowTwo, arr[i+1])
				copy(rowThree, arr[i+2])
				middleNum := rowTwo[j+1]
				hourGlassArray = append(rowOne[j:rowSlice], rowThree[j:rowSlice]...)
				for _, num := range hourGlassArray {
					sum += num
				}
				sum += middleNum
				if sum > max || init {
					max = sum
					init = false
				}
			}
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32

	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	fmt.Println(arr)
	result := hourglassSum(arr)
	fmt.Println(result)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
