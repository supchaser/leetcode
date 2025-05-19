package coderun

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var monthToZeller = map[string]int{
	"January":   13,
	"February":  14,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

var numDayToStr = map[int]string{
	0: "Saturday",
	1: "Sunday",
	2: "Monday",
	3: "Tuesday",
	4: "Wednesday",
	5: "Thursday",
	6: "Friday",
}

func Func32back1() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}

		var day, year int
		var monthStr string

		fmt.Sscanf(line, "%d %s %d", &day, &monthStr, &year)
		month, ok := monthToZeller[monthStr]
		if !ok {
			return
		}

		adjustedYear := year
		if month > 12 {
			adjustedYear--
		}

		century := adjustedYear / 100
		yearOfCentury := adjustedYear % 100
		h := (day + (13*(month+1))/5 + yearOfCentury + yearOfCentury/4 + century/4 + 5*century) % 7
		if h < 0 {
			h += 7
		}
		result, ok := numDayToStr[h]
		if !ok {
			continue
		}

		fmt.Println(result)
	}
}
