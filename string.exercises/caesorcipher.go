package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"unicode"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32, m int) string {
	alphabet := "zyxwvutsrqponmlkjihgfedcba"
	r := []rune(alphabet)
	n := len(r)
	rot := make([]rune, n)
	encryption := make(map[string]string)
	var cipher string = ""
	for i := 0; i < n; i++ {
		index := (i+int(k))%n
		rot[index] = r[i] 
		lowerAlpha := string(r[index])
		lowerAlphaRotated := string(rot[index])
		upperAlpha := strings.ToUpper(string(r[index]))
		upperAlphaRotated := strings.ToUpper(string(rot[index]))
		encryption[lowerAlpha] = lowerAlphaRotated
		encryption[upperAlpha] = upperAlphaRotated
	}
	for _,c := range s {
		value := string(c)
		if unicode.IsLetter(c) {
			value = encryption[string(c)]
		}
		cipher = cipher + value
	}
	return cipher
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int(nTemp)

    s := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k, n)

    fmt.Fprintf(writer, "%s\n", result)

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
