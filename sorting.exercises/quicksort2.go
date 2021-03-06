package main
/*
    quicksort algorithm that uses the first index as the pivot
*/

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func swap(a *int32, b *int32) {
    var temp int32 = *a
    *a = *b
    *b = temp
}

func partition(arr []int32, low int, high int) int {
	pivot := arr[low]
	var i int = low+1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
            //swap
            swap(&arr[i], &arr[j])
            i++
		}
	}

    swap(&arr[low], &arr[i-1])
	return i-1
}

// Complete the quickSort function below.
func quickSort(arr []int32, low int, high int) {
	if low < high {

		//Select pivot position
		pi := partition(arr, low, high)

		//sort elements on the left of the pivot
		quickSort(arr, low, pi - 1)
		quickSort(arr, pi + 1, high)
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

    quickSort(arr, 0, int(n)-1)

    for _, resultItem := range arr {
        fmt.Printf("%d ", resultItem)
    }

    fmt.Printf("\n")
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
