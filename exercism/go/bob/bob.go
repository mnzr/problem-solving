package bob

// package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isNumOnly(remark string) bool {
	reg, _ := regexp.Compile("[^0-9]+")
	processedString := reg.ReplaceAllString(remark, "")
	if processedString == "" {
		return false
	}
	return true
}

func isQuestion(remark string) bool {
	if string(remark[len(remark)-1]) == "?" {
		return true
	}
	return false
}

func isExclamation(remark string) bool {
	if string(remark[len(remark)-1]) == "!" {
		return true
	}
	return false
}

func isYelling(remark string) bool {
	if !isNumOnly(remark) && remark == strings.ToUpper(remark) {
		return true
	}
	return false
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	// if you address him without actually saying anything.
	if remark == "" {
		return "Fine. Be that way!"
	}

	if isQuestion(remark) {
		if isYelling(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if isYelling(remark) {
		return "Whoa, chill out!"
	}

	if isNumOnly(remark) {
		return "Sure."
	}
	return "Whatever."
}

func main() {
	fmt.Println(Hey("Ki hoise?"))
	fmt.Println(Hey("KI HOISE!"))
	fmt.Println(Hey("Ki hoise!"))
	fmt.Println(Hey("\t\t\t\t\t\t\t\t\t\t"))
}
