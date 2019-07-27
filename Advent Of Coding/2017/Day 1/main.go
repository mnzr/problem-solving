package main

import (
	"fmt"
)

func findSequence(numbers []int) int {
	var multipleOccurances []int
	sum := 0
	previousNumber := numbers[0]
	if previousNumber == numbers[len(numbers)-1] {
		ok := false
		for _, s := range multipleOccurances {
			if s == previousNumber {
				ok = true
			}
		}
		if ok != true {
			multipleOccurances = append(multipleOccurances, previousNumber)
			sum += previousNumber
		}
	}
	for _, currentNumber := range numbers[1:] {
		if previousNumber == currentNumber {
			ok := false
			for s := range multipleOccurances {
				if s == previousNumber {
					ok = true
				}
			}
			if ok != true {
				multipleOccurances = append(multipleOccurances, previousNumber)
				sum += previousNumber
			}
		} else {
			previousNumber = currentNumber
		}
	}
	return sum
}

func takeInput() []int {
	print("enter a number: ")
	var inputNumber int
	fmt.Scan(&inputNumber)
	const factor = 10
	var inputSequence []int
	for inputNumber > 0 {
		inputSequence = append(inputSequence, inputNumber%factor)
		inputNumber = inputNumber / factor
	}
	return inputSequence
}

func main() {
	fmt.Println(findSequence(takeInput()))
}
