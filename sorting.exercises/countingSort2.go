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
func countSort(arr [][]string) {

	size := len(arr)
	firstHalf := size / 2

	var max int

	// find the max [int(string) string]
	for idx, tup := range arr {
		j, _ := strconv.Atoi(tup[0])
		if j > max {
			max = j
		}

		// The first half of the strings encountered in the inputs are to be replaced with the character "-".
		if idx < firstHalf {
			tup = arr[idx]
			tup[1] = "-"
			arr[idx] = tup
		}
	}

	var output = make([]string, max+1)
	for i := 0; i < size; i++ {
		index, _ := strconv.Atoi(arr[i][0])
		output[index] += arr[i][1]
		output[index] += " "
	}

	fmt.Println(strings.Join(output, ""))
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
