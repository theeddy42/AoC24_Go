package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func removeElement(array []string, index int) []string {
	var intList []string
	for i := 0; i < len(array); i++ {
		if i != index {
			intList = append(intList, array[i])
		}
	}
	return intList
}

func convertStringArrayToInt(array []string) []int {
	var intList []int
	for _, val := range array {
		read_int, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		intList = append(intList, read_int)
	}
	return intList
}

func checkLevelSafe(level []string) bool {
	var intList = convertStringArrayToInt(level)
	var retVal = true
	var direction = 0 // 1 for up, 2 for down 0 default
	for index, val := range intList {
		if index == 0 {
			// Skip first entry because we want to check with n-1
			continue
		}

		// get initial direction
		if direction == 0 && val < intList[index-1] {
			direction = 2 //down
		} else if direction == 0 && val > intList[index-1] {
			direction = 1 //up
		}

		if (direction == 1 && val < intList[index-1]) || (direction == 2 && val > intList[index-1]) {
			retVal = false
		}

		// Must be decreasing and may only decrease by 1 to 3
		if !(math.Abs(float64(intList[index-1]-val)) <= 3) || intList[index-1]-val == 0 {
			retVal = false
		}

	}
	return retVal
}
func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	/*
	 Split by \r\n to get each lines.
	*/
	input := strings.Split(string(input_bytes), "\r\n")

	var total_safe_levels_part1 = 0
	var total_safe_levels_part2 = 0
	for line := 0; line < len(input); line++ {
		// Get Fields from line
		input_fields := strings.Fields(input[line])
		// check if length of input_fields is greater 0 because strings.Fields may return an empty array
		isLevelSafe := checkLevelSafe(input_fields)
		if len(input_fields) > 0 && isLevelSafe {
			total_safe_levels_part1++
			total_safe_levels_part2++
		}
		if len(input_fields) > 0 && !isLevelSafe {
			for index, _ := range input_fields {
				new_line := removeElement(input_fields, index)
				isNewLevelSafe := checkLevelSafe(new_line)
				if len(new_line) > 0 && isNewLevelSafe {
					total_safe_levels_part2++
					fmt.Println(new_line)
					break
				}
			}
		}

	}
	fmt.Println("Number of safe levels:", total_safe_levels_part1)
	fmt.Println("Number of safe levels with tolerance:", total_safe_levels_part2)

}
