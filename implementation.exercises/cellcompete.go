package main

/*
	There is a colony of 8 cells arranged in a straight line where each day every cell competes with its adjacent cells(neighbour).
	Each day, for each cell, if its neighbours are both active or both inactive, the cell becomes inactive the next day,. otherwise itbecomes active the next day.

	Assumptions: The two cells on the ends have single adjacent cell, so the other adjacent cell can be assumsed to be always inactive. Even after updating the cell state. 
	consider its pervious state for updating the state of other cells. Update the cell informationof allcells simultaneously.

	Write a fuction cellCompete which takes takes one 8 element array of integers cells representing the current state of 8 cells and one integer days representing te number of days to simulate. 
	An integer value of 1 represents an active cell and value of 0 represents an inactive cell.
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func cellCompete(cells []int32, days int) []int32 {
	n := len(cells)
	temp := make([]int32, n)
	copy(temp, cells)
	for i := 0; i < days; i++ {
		for j := 1; j < n-1; j++ {
			if cells[j-1] == cells[j+1] {
				temp[j] = 0
			} else {
				temp[j] = 1
			}
		}
		if cells[1] == 0 {
			temp[0] = 0 
		} else {
			temp[0] = 1
		}
		if cells[n-2] == 0 {
			temp[n-1] = 0 
		} else {
			temp[n-1] = 1
		}
		copy(cells, temp)
	}
	return cells
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	l := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(l[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	cellTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var cells []int32

	for i := 0; i < int(n); i++ {
		cellItemTemp, err := strconv.ParseInt(cellTemp[i], 10, 64)
		checkError(err)
		cellItem := int32(cellItemTemp)
		cells = append(cells, cellItem)
	}

	dTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	d := int(dTemp)

	result := cellCompete(cells, d)
	fmt.Println(result)
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

