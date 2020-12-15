package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the pageCount function below.
 */
func pageCount(n int32, p int32) int32 {
	var qIndex int32 = p - 1
	var bookEnd int32 = n - 1
	var fromBegginingCount int32 = 0
	var fromEndCount int32 = 0
	var i int32

	if qIndex%2 != 0 {
		qIndex++
	}

	for i = 0; i < qIndex; i += 2 {
		fromBegginingCount++
	}

	for i = bookEnd; i > qIndex; i -= 2 {
		fromEndCount++
	}

	if fromBegginingCount < fromEndCount {
		return fromBegginingCount
	}
	return fromEndCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

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
