package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the theLoveLetterMystery function below.
func theLoveLetterMystery(s string) int32 {
	var reductions int32 = 0
	var lastIndex = len(s)-1
	for i := 0; i < lastIndex-i; i++ {
		if s[i] > s[lastIndex-i] {
			reductions += int32(s[i]-s[lastIndex-i])
		}
		if s[i] < s[lastIndex-i] {
			reductions += int32(s[lastIndex-i]-s[i])
		}
	}
	return reductions
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s := readLine(reader)

        result := theLoveLetterMystery(s)

        fmt.Fprintf(writer, "%d\n", result)
    }

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