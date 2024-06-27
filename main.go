package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: your-program \"*/15 0 1,15 * 1-5 /usr/bin/find\"")
		return
	}

	cronExpression := os.Args[1]
	fields := strings.Fields(cronExpression)
	if len(fields) != 6 {
		fmt.Println("Invalid cron expression. Please provide exactly five time fields and a command.")
		return
	}

	minuteField := fields[0]
	hourField := fields[1]
	dayOfMonthField := fields[2]
	monthField := fields[3]
	dayOfWeekField := fields[4]
	command := fields[5]

	fmt.Printf("minute        %s\n", expandField(minuteField, 0, 59))
	fmt.Printf("hour          %s\n", expandField(hourField, 0, 23))
	fmt.Printf("day of month  %s\n", expandField(dayOfMonthField, 1, 31))
	fmt.Printf("month         %s\n", expandField(monthField, 1, 12))
	fmt.Printf("day of week   %s\n", expandField(dayOfWeekField, 0, 6))
	fmt.Printf("command       %s\n", command)
}

func expandField(field string, min int, max int) string {
	if field == "*" {
		return rangeToString(min, max, 1)
	}

	if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		var result []string
		for _, part := range parts {
			result = append(result, expandSingleField(part, min, max))
		}
		return strings.Join(result, " ")
	}

	return expandSingleField(field, min, max)
}

func expandSingleField(field string, min int, max int) string {
	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		step, _ := strconv.Atoi(parts[1])
		return rangeToString(min, max, step)
	}

	if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		return rangeToString(start, end, 1)
	}

	num, _ := strconv.Atoi(field)
	return strconv.Itoa(num)
}

func rangeToString(start int, end int, step int) string {
	var result []string
	for i := start; i <= end; i += step {
		result = append(result, strconv.Itoa(i))
	}
	return strings.Join(result, " ")
}
