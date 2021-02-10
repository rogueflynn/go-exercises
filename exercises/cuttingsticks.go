package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) []int32 {
    var sticksCut []int32
    sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
    var sorted = make([]int32, len(arr))
    copy(sorted, arr)
    for len(sorted) > 0 {
       var cuts []int32 
       num := sorted[0]
       for i := 0; i < len(sorted); i++ {
          sorted[i] -= num 
          if sorted[i] > 0 {
              cuts = append(cuts, sorted[i])
          }
       }
       sticksCut = append(sticksCut, int32(len(sorted)))
       sorted = cuts
    }
    return sticksCut
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

    result := cutTheSticks(arr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
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
