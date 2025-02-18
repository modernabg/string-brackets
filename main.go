package main

import (
	"fmt"
)

func isValidBracketSequence(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	openingBrackets := map[rune]bool{'(': true, '[': true, '{': true}

	for _, char := range s {
		if openingBrackets[char] {
			stack = append(stack, char)
		} else if matchingChar, exists := bracketMap[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != matchingChar {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(isValidBracketSequence("({}[]){}"))         // True
	fmt.Println(isValidBracketSequence("([([([])])])"))     // True
	fmt.Println(isValidBracketSequence("(()]"))              // False
	fmt.Println(isValidBracketSequence(")()("))              // False
	fmt.Println(isValidBracketSequence("{}("))               // False
	fmt.Println(isValidBracketSequence("{(})"))              // False
}
