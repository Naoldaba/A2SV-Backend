package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Palindrome() {
	fmt.Println("Welcome! Please enter any string.")

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	str := reader.Text()

	if str == "" {
		fmt.Println("Input cannot be empty.")
		return
	}

	input := strings.ToLower(str)

	var cleanedRune []rune
	for _, char := range input {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			cleanedRune = append(cleanedRune, char)
		}
	}

	if len(cleanedRune) == 0 {
		fmt.Println("Input must contain letters or digits.")
		return
	}

	n := len(cleanedRune)
	for i := 0; i < n/2; i++ {
		if cleanedRune[i] != cleanedRune[n-i-1] {
			fmt.Println("Not a palindrome")
			return
		}
	}
	fmt.Println("It is a palindrome")
}

func main() {
	Palindrome()
}
