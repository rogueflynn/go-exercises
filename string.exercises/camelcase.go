package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
	"unicode"
)

// Complete the camelcase function below.
func camelcase(s string) int32 {
	var word string = ""
	var words []string
	for _,r := range s {
		if unicode.IsUpper(r) {
			words = append(words, word)
			word = ""
		}
		word = word + string(r)
	}
	words = append(words, word) // append last word
	return int32(len(words))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := camelcase(s)

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
