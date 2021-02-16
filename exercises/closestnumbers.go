package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"sort"
	"math"
)

// Complete the closestNumbers function below.
func closestNumbers(arr []int32) []int32 {
	var max int = 0
	for _,item := range arr {
		num := int(math.Abs(float64(item)))
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
	var pairs = make([][]int32, max)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j]})	
	var min int = 0
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			diff := int(math.Abs(float64(arr[i] - arr[i+1])))
			if diff <= min  || min == 0 {
				min = diff
				pairs[min] = append(pairs[min], arr[i])
				pairs[min] = append(pairs[min], arr[i+1])
			}
		}
	}
	return pairs[min]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := closestNumbers(arr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
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
