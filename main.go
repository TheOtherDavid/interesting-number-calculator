package main

import (
	"fmt"
	"math"
)

type NumberInfo struct {
	Value       int
	Interesting bool
}

// strictness defines how many digits in a row will trigger the test
// for a strictness of 3, a straight could be 000123
// for a strictness of 4, a straight would require 001234
const straightStrictness = 3
const repeatStrictness = 3

func main() {
	//let's start with the number of digits
	//eventually this will be passed in as an arg
	//start with 3 for debugging faster
	digits := 6

	interestingCount := 0
	formatString := fmt.Sprintf("%%0%dd", digits)
	var numbers []NumberInfo

	maxValue := int(math.Pow10(digits)) - 1

	for i := 0; i <= maxValue; i++ {
		//now we do our interestingness tests
		isInteresting := false
		numStr := fmt.Sprintf(formatString, i)
		if isNumberInteresting(numStr) {
			isInteresting = true
			interestingCount++
			//fmt.Println(numStr, isInteresting)
		}
		numbers = append(numbers, NumberInfo{Value: i, Interesting: isInteresting})

	}
	totalNumbers := float64(maxValue + 1)
	interestingFloat := float64(interestingCount)
	//fmt.Printf("%d out of %f numbers are interesting\n", interestingCount, totalNumbers)
	percentInteresting := (interestingFloat / totalNumbers) * 100
	fmt.Printf("%.3f percent of %d digit numbers are interesting", percentInteresting, digits)
}

func isNumberInteresting(numStr string) bool {
	//Palindrome test
	if numStr == reverseString(numStr) {
		return true
	}
	//Straight test
	if hasStraight(numStr) {
		return true
	}
	//Repeat test
	if hasConsecutiveRepeat(numStr) {
		return true
	}
	return false
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func hasStraight(numStr string) bool {
	for i := 0; i <= len(numStr)-straightStrictness; i++ {
		ascending := true
		descending := true

		for j := 0; j < straightStrictness-1; j++ {
			if numStr[i+j+1] != numStr[i+j]+1 {
				ascending = false
			}
			if numStr[i+j+1] != numStr[i+j]-1 {
				descending = false
			}
			if !ascending && !descending {
				break
			}
		}

		if ascending || descending {
			return true
		}
	}
	return false
}

func hasConsecutiveRepeat(numStr string) bool {
	for i := 0; i <= len(numStr)-repeatStrictness; i++ {
		isRepeat := true
		for j := 1; j < repeatStrictness; j++ {
			if numStr[i+j] != numStr[i] {
				isRepeat = false
				break
			}
		}
		if isRepeat {
			return true
		}
	}
	return false
}
