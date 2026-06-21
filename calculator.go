// Build a cli calculator in golang/
// Let's go

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	calculatorRunning := true
	scanner := bufio.NewScanner(os.Stdin)
	for calculatorRunning {
		fmt.Print("Enter digits to calculate: e.g 10 + 3:\n")
		scanner.Scan()
		text := scanner.Text()
		digit := find_digit(text)
		operator := find_operator(text)

		re := regexp.MustCompile(`[^0-9+\-*/\s]`)
		text_regexp := re.FindAllString(text, -1)

		switch {
		case len(digit) == 0 || len(operator) == 0 || len(text_regexp) >= 1:
			fmt.Println("Input must be a digit and operator")
			continue
		}

		if len(digit) != len(operator)+1 {
			fmt.Println("Invalid input operation")
			continue
		} else {
			result := calculator(digit, operator)
			fmt.Println("Result is:", result)
		}
	}

}

func find_digit(digits string) []string {
	re := regexp.MustCompile(`\d+`)

	digitMatched := re.FindAllString(digits, -1)

	// fmt.Println(digitMatched)
	return digitMatched
}

func find_operator(operators string) []string {
	re := regexp.MustCompile(`[+\-*/]`)

	operatorMatched := re.FindAllString(operators, -1)
	// fmt.Println(operatorMatched)

	return operatorMatched
}

func calculator(digits []string, operator []string) int {
	num_slice := convert_to_num(digits)
	var output int
	for _, s := range operator {
		if len(num_slice) >= 2 {

			switch s {
			case "+":
				output = num_slice[0] + num_slice[1]
				num_slice = num_slice[2:]
				num_slice = append([]int{output}, num_slice...)
			case "-":
				output = num_slice[0] - num_slice[1]
				num_slice = num_slice[2:]
				num_slice = append([]int{output}, num_slice...)
			case "*":
				output = num_slice[0] * num_slice[1]
				num_slice = num_slice[2:]
				num_slice = append([]int{output}, num_slice...)
			case "/":
				output = num_slice[0] / num_slice[1]
				num_slice = num_slice[2:]
				num_slice = append([]int{output}, num_slice...)
			}

		}
	}
	return output
}


func convert_to_num(digits []string) []int {
	nums := make([]int, 0, len(digits))

	for _, s := range digits {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid number:", s)
			continue
		}
		nums = append(nums, n)
	}

	return nums
}
