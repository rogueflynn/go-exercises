package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "regexp"
)

// Complete the minimumNumber function below.
func minimumNumber(n int32, password string) int32 {
    // Return the minimum number of characters to make the password strong
    var charsNeeded int32 = 0
    var passwordLength int32 = int32(len(password))
    if matched,_ := regexp.MatchString("[a-z]", password); !matched {
        charsNeeded++
    }
    if matched,_ := regexp.MatchString("[A-Z]", password); !matched {
        charsNeeded++
    }
    if matched,_ := regexp.MatchString("[0-9]", password); !matched {
        charsNeeded++
    }
    if matched,_ := regexp.MatchString("[!@#$%^&*()\\-+]", password); !matched {
        charsNeeded++
        fmt.Println("not matched")
    }

    passwordLength = passwordLength + charsNeeded
    if passwordLength < 6 {
        diff :=  6 - passwordLength
        charsNeeded = charsNeeded + diff
    }

    return charsNeeded;
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

    password := readLine(reader)

    answer := minimumNumber(n, password)

    fmt.Fprintf(writer, "%d\n", answer)

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
