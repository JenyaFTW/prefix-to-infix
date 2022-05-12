package main

import (
	"strings"
)

func prefixEquationToInfix(prefix string) string {
	prefixArray := strings.Split(prefix, " ")

	stack := make([]string, 0)

	var operatorCount int
	var operandCount int
	for _, element := range prefixArray {
		if element == "+" || element == "-" || element == "*" || element == "/" || element == "^" {
			operatorCount++
		} else {
			operandCount++
		}
	}

	if operatorCount != operandCount-1 {
		return "Invalid prefix equation"
	}

	for i := len(prefixArray) - 1; i >= 0; i-- {
		if prefixArray[i] != "+" && prefixArray[i] != "-" && prefixArray[i] != "*" && prefixArray[i] != "/" && prefixArray[i] != "^" {
			stack = append(stack, prefixArray[i])
		} else {
			if len(stack) >= 2 {
				num1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				num2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, "("+num1+prefixArray[i]+num2+")")
			} else {
				return "Invalid prefix equation"
			}
		}
	}

	return stack[0]
}
