package timeparse

//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg, input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v, %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {

	timeString := strings.Split(input, ":")

	if len(timeString) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time part component", input}
	}
	hour, err := strconv.Atoi(timeString[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", hour), input}
	}
	minute, err := strconv.Atoi(timeString[1])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", minute), input}
	}
	second, err := strconv.Atoi(timeString[2])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second: %v", second), input}
	}

	if hour > 23 || hour < 0 {
		return Time{}, &TimeParseError{"hour not in range : 0 <= %v => ", fmt.Sprintf("%v", hour)}
	}
	if minute > 59 || minute < 0 {
		return Time{}, &TimeParseError{"hour not in range : 0 <= %v => ", fmt.Sprintf("%v", minute)}
	}
	if second > 59 || second < 0 {
		return Time{}, &TimeParseError{"hour not in range : 0 <= %v => ", fmt.Sprintf("%v", second)}
	}
	return Time{hour, minute, second}, nil
}
