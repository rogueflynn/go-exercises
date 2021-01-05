package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
func printArray(arr []int32) {
    for _,element := range arr {
       fmt.Printf("%d ", element) 
    }
    fmt.Printf("\n") 
}

// Complete the insertionSort1 function below.
func insertionSort1(n int32, arr []int32) {
    var i int32 = 0
    temp := arr[n-1]
    for i = n-1; i > -1; i-- {
        if i != 0 && arr[i-1] > temp {
            arr[i] = arr[i-1]
            printArray(arr)
        } else if i != 0 && temp > arr[i-1] {
            arr[i] = temp
			printArray(arr)
			break
        } else if i == 0 {
            arr[i] = temp
			printArray(arr)
		}     
	}
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

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

    insertionSort1(n, arr)
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
