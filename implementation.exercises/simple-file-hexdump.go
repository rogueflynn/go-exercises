/*
	Simple script that performs a hexdump of a file
	using go's hex package.

	Input: Any data file.  (not tested on large binaries)
	       Path must be relative to the where the script is.

	Output: Hex dump with similar results as hexdump -C

	Usage:
		go run simple-file-hexdump.go <file>
*/
package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		file, err := os.Open("./" + filename)
		check(err)
		defer file.Close()
		fileinfo, err := file.Stat()
		check(err)
		buffer := make([]byte, fileinfo.Size())
		file.Read(buffer)
		fmt.Println(hex.Dump(buffer))

	} else {
		fmt.Println("Please enter a file to parse")
	}
}
