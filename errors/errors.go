package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at index %v", index))
	} else {
		return s.values[index], nil
	}
}

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func Parse(time string) (Time, error) {

	components := strings.Split(time, ":")

	if len(components) != 3 {
		return Time{}, &TimeParseError{msg: "Invalid number of time components", input: time}
	} else {
		hour, err := strconv.Atoi(components[0])

		if err != nil {
			return Time{}, &TimeParseError{msg: fmt.Sprintf("Error parsing hour: %v", err), input: time}

		}

		minute, err := strconv.Atoi(components[1])

		if err != nil {
			return Time{}, &TimeParseError{msg: fmt.Sprintf("Error parsing minute: %v", err), input: time}

		}

		second, err := strconv.Atoi(components[2])

		if err != nil {
			return Time{}, &TimeParseError{msg: fmt.Sprintf("Error parsing second: %v", err), input: time}

		}

		return Time{hour: hour, minute: minute, second: second}, nil
	}
}

func main() {

	stuff := Stuff{}

	value, err := stuff.Get(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}

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

	table := []struct {
		time string
	}{
		{"19:00:12"},
		{"1:3:44"},
		{"bad"},
		{"1:-3:44"},
		{"0:59:59"},
		{""},
		{"11:22"},
		{"aa:bb:cc"},
		{"5:23:"},
	}

	for _, data := range table {
		result, err := Parse(data.time)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
