package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func dist(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func countOccurences(array []int, target int) int {
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			count++
		}
	}
	return count
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Split input into two lists
	var left_list, right_list []int
	var found_left bool = false

	input := strings.Fields(string(input_bytes))
	for i := 0; i < len(input); i++ {
		number := string(input[i])
		if number != "" && !found_left {
			found_left = true
			read_int, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Error reading left byte", input[i])
			}
			left_list = append(left_list, read_int)
		} else if number != "" && found_left {
			found_left = false
			read_int, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Error reading right byte", input[i])
			}
			right_list = append(right_list, read_int)
		} else {
			// Do nothing
		}
	}
	sort.Ints(left_list)
	sort.Ints(right_list)
	total_distance := 0
	for i := 0; i < len(left_list); i++ {
		total_distance += dist(left_list[i], right_list[i])
	}
	fmt.Println("Total distance: ", total_distance)

	// part 2 - similarity count
	// Go through each left_list, find no. of occurences of number in left_list within right_list and multiply by that value
	// Sum the similarity values

	total_similarity := 0
	for i := 0; i < len(left_list); i++ {
		occurences := countOccurences(right_list, left_list[i])
		total_similarity += left_list[i] * occurences
	}
	fmt.Println("Total similarity: ", total_similarity)
}
