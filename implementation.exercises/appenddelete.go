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
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

 /*
 * Case 1: See if we can completely erase String s and append String t. If we need to waste operations 
 * to reach k operations, we can do so when String s has no characters.
 *
 * Case 2: See if we can convert String s to String t without completely erasing String s. 
 * We keep erasing charcters from String s until it becomes a prefix of String t. We then add 
 * the characters needed to turn String s into String t. If we need to waste operations to 
 * reach k operations, we can only do so in groups of 2 by doing an append and a delete.
 *
 */


func appendAndDelete(s string, t string, k int32) string {
    // Write your code here
    // Write your code here
    
    L := int32(len(s) + len(t))
    var i int

    if L <= k {
        return "Yes"
    }

    for i = 0; i < len(s) && i < len(t); i++ {
        if s[i] != t[i] {
            break
        }
    }

    minOperations := int32((len(s) - i) + (len(t) - i))

    if k >= minOperations && (k - minOperations) % 2 == 0 {
        return "Yes"
    }

    return "No"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    t := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := appendAndDelete(s, t, k)

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