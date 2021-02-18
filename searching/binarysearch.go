package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid]  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	var arr []int
	var index int
	for i := 10; i <= 30; i++ {
		arr = append(arr, i)
	}
	index = binarySearch(arr, 11)  
	fmt.Println(index) //should return 1
	index = binarySearch(arr, 15)  
	fmt.Println(index) //should return 5
	index = binarySearch(arr, 50)  
	fmt.Println(index) //should return -1
}