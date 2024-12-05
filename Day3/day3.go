package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calculateMul(found_muls []string, sum_multiplication *int) {
	just_mul, _ := regexp.Compile("[a-z\\(\\)]*")
	for _, mul := range found_muls {
		numbers := strings.Split(just_mul.ReplaceAllString(mul, ""), ",")
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error reading first number")
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error reading first number")
		}
		*sum_multiplication += x * y
	}
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	input := strings.ReplaceAll(string(input_bytes), "\r\n", "")

	sum_multiplication := 0

	mul_regex, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	found_muls := mul_regex.FindAllString(input, -1)
	calculateMul(found_muls, &sum_multiplication)

	fmt.Println("Part 1: ", sum_multiplication)
	// Part 2
	sum_multiplication = 0
	// First regex to find everything in between don't() and do()'s.
	dontdo_regex, _ := regexp.Compile("don't\\(\\)(.*?)do\\(\\)")
	// Second Regex for any don't() calls that do not end with a do()
	dont_regex, _ := regexp.Compile("don't\\(\\)(.*)")
	// Remove all substrings between don't() and do()'s
	dont_do_strings := dontdo_regex.FindAllString(input, -1)
	for _, dont_do_str := range dont_do_strings {
		input = strings.ReplaceAll(input, dont_do_str, "")

	}
	// remove last possible don't() call that may not end with another do()
	input = strings.ReplaceAll(input, dont_regex.FindString(input), "")

	found_muls = mul_regex.FindAllString(input, -1)
	calculateMul(found_muls, &sum_multiplication)

	fmt.Println("Part 2: ", sum_multiplication)
}
