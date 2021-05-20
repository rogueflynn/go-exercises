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
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

func timeInWords(h int32, m int32) string {
    // Write your code here
	var time string = ""
	hour := make(map[string]string)
	minute := make(map[string]string)
	hour["1"] = "one"
	hour["2"] = "two"
	hour["3"] = "three"
	hour["4"] = "four"
	hour["5"] = "five"
	hour["6"] = "six"
	hour["7"] = "seven"
	hour["8"] = "eight"
	hour["9"] = "nine"
	hour["10"] = "ten"
	hour["11"] = "eleven"
	hour["12"] = "twelve"

	minute["1"]	 = "one minute"
	minute["2"]	 = "two minutes"
	minute["3"]	 = "three minutes"
	minute["4"]	 = "four minutes"
	minute["5"]	 = "five minutes"
	minute["6"]	 = "six minutes"
	minute["7"]	 = "seven minutes"
	minute["8"]	 = "eight minutes"
	minute["9"]	 = "nine minutes"
	minute["10"] = "ten minutes"
	minute["11"] = "eleven minutes"
	minute["12"] = "twelve minutes"
	minute["13"] = "thirteen minutes"
	minute["14"] = "fourteen minutes"
	minute["15"] = "quarter"
	minute["16"] = "sixteen minutes"
	minute["17"] = "seventeen minutes"
	minute["18"] = "eighteen minutes"
	minute["19"] = "nineteen minutes"
	minute["20"] = "twenty minutes"
	minute["21"] = "twenty one minutes"
	minute["22"] = "twenty two minutes"
	minute["23"] = "twenty three minutes"
	minute["24"] = "twenty four minutes"
	minute["25"] = "twenty five minutes"
	minute["26"] = "twenty six minutes"
	minute["27"] = "twenty seven minutes"
	minute["28"] = "twenty eight minutes"
	minute["29"] = "twenty nine minutes"
	minute["30"] = "half"
	minute["31"] = "thirty one minutes"
	minute["32"] = "thirty two minutes"
	minute["33"] = "thirty three minutes"
	minute["34"] = "thirty four minutes"
	minute["35"] = "thirty five minutes"
	minute["36"] = "thirty six minutes"
	minute["37"] = "thirty seven minutes"
	minute["38"] = "thirty eight minutes"
	minute["39"] = "thirty nine minutes"
	minute["40"] = "forty minutes"
	minute["41"] = "forty one minutes"
	minute["42"] = "forty two minutes"
	minute["43"] = "forty three minutes"
	minute["44"] = "forty four minutes"
	minute["45"] = "forty five minutes"
	minute["46"] = "forty six minutes"
	minute["47"] = "forty seven minutes"
	minute["48"] = "forty eight minutes"
	minute["49"] = "forty nine minutes"
	minute["50"] = "fifty minutes"
	minute["51"] = "fifty one minutes"
	minute["52"] = "fifty two minutes"
	minute["53"] = "fifty three minutes"
	minute["54"] = "fifty four minutes"
	minute["55"] = "fifty five minutes"
	minute["56"] = "fifty six minutes"
	minute["57"] = "fifty seven minutes"
	minute["58"] = "fifty eight minutes"
	minute["59"] = "fifty nine minutes"

	hKey := strconv.FormatInt(int64(h), 10)
	mKey := strconv.FormatInt(int64(m), 10)
	if m == 0 {
		time = hour[hKey] + " o' clock"
	}
	if 1 <= m && m <= 30 {
		time = minute[mKey] + " past " + hour[hKey] 
	}
	if 30 < m {
		nextHour := h+1
		if nextHour > 12 {
			nextHour = 1
		}

		hourKey := strconv.FormatInt(int64(nextHour), 10)
		minuteKey := strconv.FormatInt(int64(60-m), 10)
		time = minute[minuteKey] + " to " + hour[hourKey] 
	}

	return time
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    h := int32(hTemp)

    mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    m := int32(mTemp)

    result := timeInWords(h, m)

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