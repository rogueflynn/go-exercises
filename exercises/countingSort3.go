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

// Complete the countSort function below.
// This is not my solution.
// Original solution author.
// https://www.hackerrank.com/challenges/countingsort4/forum/comments/872043
// All my original solution would not pass test case 5
func countSort(arr [][]string) {
	var max int
	size := len(arr)

	// find the max [int(string) string]
	for idx, tup := range arr {
		i, _ := strconv.Atoi(tup[0])
		if i > max {
			max = i
		}

		// The first half of the strings encountered in the inputs are to be replaced with the character "-".
		if idx < size/2 {
			tup = arr[idx]
			tup[1] = "-"
			arr[idx] = tup
		}
	}

	// store the count of each elements in count array
	count := make([][]string, max+1)
	for _, tup := range arr {
		i, _ := strconv.Atoi(tup[0])

		var currCount int // = 0

		if count[i] == nil {

			count[i] = []string{"1", tup[1]}

		} else {

			currCount, _ = strconv.Atoi(count[i][0])
			count[i][0] = strconv.Itoa(currCount + 1)

			count[i] = append(count[i], tup[1])

		}
	}

	// plaining slices and removing the inside index
	var out []string
	for _, val := range count {
		if val != nil {
			out = append(out, val[1:]...)
		}
	}
	fmt.Println(strings.Join(out, " "))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]string
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []string
		for _, arrRowItem := range arrRowTemp {
			arrRow = append(arrRow, arrRowItem)
		}

		if len(arrRow) != 2 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	countSort(arr)
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
