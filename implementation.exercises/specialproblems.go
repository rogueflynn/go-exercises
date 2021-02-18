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

// Complete the workbook function below.
func workbook(n int32, k int32, arr []int32) int32 {
	var totalPages int = 1 
	var special int32 = 0
	for i := 0; i < int(n); i++ {
		pages := int(math.Ceil(float64(arr[i])/float64(k)))	
		startingPage := totalPages
		// endPage := pages + startingPage - 1
		for j := 1; j <= int(arr[i]); j++ {
			if j == startingPage {
				special++
			}
			if( j % int(k) == 0) {
				startingPage++
			}
		} 
		totalPages += pages
	}
	return special
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := workbook(n, k, arr)

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
