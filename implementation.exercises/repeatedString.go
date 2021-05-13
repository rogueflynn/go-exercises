package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"math"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
    // Write your code here
	L := int64(len(s))
	var count int64 = 0
	var i int64

	val := int64(math.Floor(float64(float64(n)/float64(L))))
    prod := int64(L * int64(val))
    diff := int64(n - prod)
	for i = 0; i < L; i++ {
		if string(s[i]) == "a" {
			count++
		}
	}

    var aCount int64  = count * val

    if diff <= L {
        for i = 0; i < diff; i++ {
            if string(s[i]) == "a" {
                aCount++
            }
        }
    }

	return aCount
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

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
