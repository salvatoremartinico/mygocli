package mygocli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Saluta restituisce un saluto per la persona specificata.
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

func ReadString(scanner *bufio.Scanner, label string) string {
	fmt.Println(label)
	return strings.TrimSpace(scanner.Text())
}

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
