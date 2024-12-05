package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	x int
	y int
}
type update struct {
	update []int
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

func convertStringToArray(str string) []string {
	var list []string
	str_split := strings.Split(str, ",")
	for _, val := range str_split {
		list = append(list, string(val))
	}
	return list
}

func indexFromIntArray(arr []int, element int) int {
	retVal := -1
	for index, value := range arr {
		if value == element {
			retVal = index
		}
	}
	return retVal
}

/*
Iterate through the rules and check if the index of x is greater than the index of y.
In that case the update is not valid as the index of x must always be smaller than the index of y.
*/
func applyRules(update []int, rules []rule) bool {
	valid_update := true
	for _, rule := range rules {
		x_index := indexFromIntArray(update, rule.x)
		y_index := indexFromIntArray(update, rule.y)
		if x_index != -1 && y_index != -1 && x_index > y_index {
			valid_update = false
			break
		}
	}
	return valid_update
}

/*
Make invalid blocks valid by switching invalid indexes in-place with each other and rechecking for a valid state.
*/
func applyRulesAndSwitch(update *[]int, rules []rule) bool {
	valid_update := true
	for !applyRules(*update, rules) {
		for _, rule := range rules {
			x_index := indexFromIntArray(*update, rule.x)
			y_index := indexFromIntArray(*update, rule.y)
			if x_index != -1 && y_index != -1 && x_index > y_index {
				temp := (*update)[x_index]
				(*update)[x_index] = (*update)[y_index]
				(*update)[y_index] = temp
			}
		}
	}
	return valid_update
}

func main() {
	input_bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var rules []rule
	var valid_updates []update

	input := strings.Split(string(input_bytes), "\r\n")

	for _, line := range input {
		if strings.Contains(line, "|") {
			line_split := strings.Split(line, "|")
			x_int, _ := strconv.Atoi(line_split[0])
			y_int, _ := strconv.Atoi(line_split[1])
			rules = append(rules, rule{x: x_int, y: y_int})
		}
		if strings.Contains(line, ",") {
			update_str_array := convertStringToArray(line)
			update_int_array := convertStringArrayToInt(update_str_array)
			valid_updates = append(valid_updates, update{update: update_int_array})
		}
	}

	var invalid_updates []update

	sum_middle_page_number := 0
	for _, update := range valid_updates {
		valid_update := applyRules(update.update, rules)
		if valid_update {
			update_len := len(update.update)
			middle_index := math.Ceil(float64(update_len)/2) - 1
			sum_middle_page_number += update.update[int(middle_index)]
		} else {
			invalid_updates = append(invalid_updates, update)
		}
	}
	fmt.Println("Sum of middle pages: ", sum_middle_page_number)

	// ########### Part 2
	sum_middle_page_number = 0

	for _, update := range invalid_updates {
		valid_update := applyRulesAndSwitch(&update.update, rules)
		if valid_update {
			update_len := len(update.update)
			middle_index := math.Ceil(float64(update_len)/2) - 1
			sum_middle_page_number += update.update[int(middle_index)]
		} else {
			invalid_updates = append(invalid_updates, update)
		}
	}
	fmt.Println("Sum of middle pages invalid updates: ", sum_middle_page_number)
}
