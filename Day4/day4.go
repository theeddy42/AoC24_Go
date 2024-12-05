package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input_bytes, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(input_bytes), "\r\n")
	xmas, _ := regexp.Compile("(XMAS)")
	xmas_backwards, _ := regexp.Compile("(XMAS)")
	for _, line := range lines {
		fmt.Println(xmas.FindAllString(line, -1))
		fmt.Println(xmas_backwards.FindAllString(line, -1))
	}
}
