package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var seaLevel int32 = 1
	var downSteps int32 = 0
	var upSteps int32 = 0
	var valleysTraversed int32 = 0
	var belowSealevel bool = false
	var aboveSealevel bool = false

	for _, c := range path {
		if string(c) == "D" {
			seaLevel--
			downSteps++
			upSteps = 0
		}
		if string(c) == "U" {
			seaLevel++
			upSteps++
			downSteps = 0
		}
		if downSteps > 0 && seaLevel < 1 && !belowSealevel {
			belowSealevel = true
			aboveSealevel = false
			valleysTraversed++
		}
		if upSteps > 0 && seaLevel > 1 && !aboveSealevel {
			belowSealevel = false
			aboveSealevel = true
		}
		if seaLevel == 1 {
			belowSealevel = false
			aboveSealevel = false
		}
	}
	return valleysTraversed
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int32(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

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
