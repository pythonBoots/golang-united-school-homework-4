package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	stringValues := strings.Fields(input)
	noSpaceString := strings.Join(stringValues, "")
	var prefix string

	if len(noSpaceString) == 0 {
		return "", errorEmptyInput
	}
	// trimming first symbol
	if string(noSpaceString[0]) == "+" {
		noSpaceString = strings.TrimLeft(noSpaceString, "+")
		prefix = "+"
	} else if string(noSpaceString[0]) == "-" {
		noSpaceString = strings.TrimLeft(noSpaceString, "-")
		prefix = "-"
	}

	// parsing numbers
	numbers := []string{}
	splittedArray := strings.Split(noSpaceString, "+")
	for _, indexOfArray := range splittedArray {
		numbers = append(numbers, strings.Split(indexOfArray, "-")...)
	}
	// generating errors
	if len(numbers) > 2 {
		return "", errorNotTwoOperands
	}

	// parsing expression
	expression := string(noSpaceString[len(numbers[0])])

	// creating numbers
	number1, err1 := strconv.Atoi(prefix + numbers[0])
	number2, err2 := strconv.Atoi(numbers[1])

	// generating errors for incorrect characters in numbers
	if err1 != nil {
		return "", fmt.Errorf("letter in first number: ", err1)
	} else if err2 != nil {
		return "", fmt.Errorf("letter in second number: ", err2)
	}

	// calculate expression
	switch expression {
	case "+":
		output = strconv.Itoa(number1 + number2)
	case "-":
		output = strconv.Itoa(number1 - number2)
	}
	return output, nil

}
