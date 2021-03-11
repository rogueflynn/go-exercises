package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the pangrams function below.
func pangrams(s string) string {
    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, " ", "")
    alphabetCount := 0
    var alphabetMap [123]int32
    for i := 97; i <= 122; i++ {
        alphabetMap[i] = 0
    }
    for _, ascii := range s {
        if alphabetMap[ascii] == 0 {
            alphabetMap[ascii] = 1
            alphabetCount++
        }
    }
    if alphabetCount == 26 {
        return "pangram"
    }
    return "not pangram"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := pangrams(s)

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