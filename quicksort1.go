package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the quickSort function below.
func quickSort(arr []int32) []int32 {
	if len(arr) < 2 {
		return arr
	} 
    var left []int32
	var right []int32 
    var sorted []int32
	pivot := arr[0]
	
	// fmt.Printf("Pivot=%d ", pivot)
    for i := 1; i < len(arr); i++ {
        if arr[i] < pivot {
			left = append(left, arr[i])
        } else {
			right = append(right, arr[i])
        }
	}
    
    left = quickSort(left)
    right = quickSort(right)
    sorted = append(sorted, left...)
    sorted = append(sorted, pivot)
    sorted = append(sorted, right...)
    sortedString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sorted)), " "), "[]")
    fmt.Println(sortedString)
    return sorted;
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

	quickSort(arr)
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
