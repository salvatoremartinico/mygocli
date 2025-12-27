// MIT License
// Copyright (c) 2025 Salvatore Martinico

package mygocli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// ForceToReadInt is a function to get an integer from user input on the terminal,
// if the user does not type a valid integer an error message (label_error ) is displayed
// and the user is prompted to enter a new number until a valid integer is entered
func ForceToReadInt(scanner *bufio.Scanner, label string, label_error string) int {
	var n64 int64
	var err error
	fmt.Println(label)
	for {
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		n64, err = strconv.ParseInt(input, 10, 0)
		if err == nil {
			break
		} else {
			fmt.Println(label_error)
		}
	}
	return int(n64)
}

// ForceToReadFloat64 is a function to get an float64 from user input on the terminal,
// if the user does not type a valid float64 an error message (label_error ) is displayed
// and the user is prompted to enter a new number until a valid float64 is entered
func ForceToReadFloat64(scanner *bufio.Scanner, label string, label_error string) float64 {
	var n64 float64
	var err error
	fmt.Println(label)
	for {
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		n64, err = strconv.ParseFloat(input, 64)
		if err == nil {
			break
		} else {
			fmt.Println(label_error)
		}
	}
	return n64
}

// ReadString is a function to get an string from user input on the terminal
func ReadString(scanner *bufio.Scanner, label string) string {
	fmt.Println(label)
	return strings.TrimSpace(scanner.Text())
}

// ReadInt is a function to get an integer from user input on the terminal
func ReadInt(scanner *bufio.Scanner, label string) (int, error) {
	var n64 int64
	var err error
	fmt.Println(label)
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	n64, err = strconv.ParseInt(input, 10, 0)
	return int(n64), err
}

// ReadFloat64 is a function to get an float64 from user input on the terminal
func ReadFloat64(scanner *bufio.Scanner, label string) (float64, error) {
	var n64 float64
	var err error
	fmt.Println(label)
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	n64, err = strconv.ParseFloat(input, 64)
	return n64, err
}
