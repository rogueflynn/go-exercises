package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

var x int32 = 0
var y int32 = 0
// Complete the insertionSort1 function below.
func insertionSort(arr []int32)  {
    var i,j, value int32
    var N = int32(len(arr))
    for i = 1; i < N; i++ {
        value = arr[i]
        j = i - 1
        
        for j >= 0 && value < arr[j] {
            arr[j+1] = arr[j]
            j = j - 1
            y++
        }
        arr[j+1] = value
    }
}

func partition(arr []int32, low int, high int) int {
	pivot := arr[high]
	var i int = low-1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			//swap
            x++
            i++
            arr[i], arr[j] = arr[j], arr[i]
		}
	}

    arr[i+1], arr[high] = arr[high], arr[i+1]
    x++
	return (i+1)
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

    var qArr = make([]int32, len(arr), len(arr))
    copy(qArr, arr)
	insertionSort(arr)
    quickSort(qArr, 0, int(n)-1)

    fmt.Printf("%d\n", y-x)
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

