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

var results int32  = 0

func checkSum(num int32, X int32, k int32, N int32) int32 {
	if X == 0 {
		results++
	}

	r := int32(math.Floor(math.Pow(float64(num), 1.0 / float64(N))))
	var i int32 = 0

	for i = k + 1; i <= r; i++ {
		var a int32 = X - int32(math.Pow(float64(i),float64(N)))

		if a >= 0 {
			checkSum(num, a, i, N)
		}
	}
	return results
}

// Complete the powerSum function below.
func powerSum(X int32, N int32) int32 {
	return checkSum(X, X, 0, N)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    XTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    X := int32(XTemp)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)

    result := powerSum(X, N)

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
