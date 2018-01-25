package bob

// package main

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	length := len(remark)
	last_char := remark[length-1]
	// fmt.Println(remark[length-1])

	if last_char == '?' {
		if remark == strings.ToUpper(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if remark == strings.ToUpper(remark) {
		return "Whoa, chill out!"
	}

	if remark == "Bob" {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

// func main() {
// 	// fmt.Println(Hey("Ki hoise?"))
// 	// fmt.Println(Hey("KI HOISE!"))
// 	// fmt.Println(Hey("Ki hoise!"))
// 	// fmt.Println(Hey("\t\t\t\t\t\t\t\t\t\t"))
// }
