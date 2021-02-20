package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the marsExploration function below.
func marsExploration(s string) int32 {
	r := []rune(s)	
	var changed int32 = 0
	for i := 0; i < len(r); i+=3 {
		if string(r[i]) != "S" {
			changed++
		}
		if string(r[i+1]) != "O" {
			changed++
		}
		if string(r[i+2]) != "S" {
			changed++
		}
	}
	return changed
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := marsExploration(s)

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
