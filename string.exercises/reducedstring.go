package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

func deleteChar(r []rune, index int) []rune {
	return append(r[0:index], r[index+1:]...)
}

// Complete the superReducedString function below.
func superReducedString(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		if string(r[i]) == string(r[i+1]) {
			r = append(r[0:i], r[i+2:]...)
			i = -1
		}
	} 
	if len(r) > 0 {
		return string(r)
	} else {
		return "Empty String"
	}
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := superReducedString(s)

    fmt.Fprintf(writer, "%s\n", result)

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

