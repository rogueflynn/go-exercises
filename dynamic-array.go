package main

/*
	Problem Solving Exercise Using go
	Parses input from a text file and runs a specific
	query based on the input from the first parameter

	Input: text file in the format:

	    2 5
	    1 0 5
	    1 1 7
	    1 0 3
	    2 1 0
	    2 1 1

	    Read from stdin

        Output:
		The values of the last answers as they are changed

	Usage:
		cat <INPUT_FILE> | go run dynamic-array.go
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dynamicArray(n int32, queries [][]int32) []int32 {
	var seqList [][]int32
	var i int32
	var length int32 = int32(len(queries))
	var lastAnswer int32 = 0
	var changedAnswers []int32

	//create sequence list
	for i = 0; i < n; i++ {
		var s []int32
		seqList = append(seqList, s)
	}

	//run queries
	for i = 0; i < length; i++ {

		var queryType = queries[i][0]
		var x = queries[i][1]
		var y = queries[i][2]
		var seqListIndex = (x ^ lastAnswer) % n
		var seq []int32 = seqList[seqListIndex]
		var size = int32(len(seq))
		var seqIndex int32 = 0

		if size != 0 {
			seqIndex = y % size
		}

		if queryType == 1 {
			seq = append(seq, y)
			seqList[seqListIndex] = seq
		} else if queryType == 2 {
			lastAnswer = seq[seqIndex]
			changedAnswers = append(changedAnswers, lastAnswer)
		}

	}

	return changedAnswers
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := dynamicArray(n, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
