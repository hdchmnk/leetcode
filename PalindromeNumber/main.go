package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	xToString := strconv.Itoa(x)

	startString := 0
	endString := len(xToString) - 1

	for startString <= endString {
		if xToString[startString] != xToString[endString] {
			return false
		}
		startString++
		endString--
	}

	return true
}
