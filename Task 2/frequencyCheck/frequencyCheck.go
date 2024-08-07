package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func frequencyCheck() {
	fmt.Println("Welcome! Please enter any string.")

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	str := reader.Text()

	str = strings.TrimSpace(str)
	if str == "" {
		fmt.Println("Input cannot be empty.")
		return
	}

	input := strings.ToLower(str)
	var cleanedInput strings.Builder

	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			_, err := cleanedInput.WriteRune(char)
			if err != nil {
				fmt.Println("Error writing rune to string builder:", err)
				return
			}
		} else {
			_, err := cleanedInput.WriteRune(' ')
			if err != nil {
				fmt.Println("Error writing rune to string builder:", err)
				return
			}
		}
	}

	words := strings.Fields(cleanedInput.String())
	if len(words) == 0 {
		fmt.Println("No valid words found in input.")
		return
	}

	wordFrequency := make(map[string]int)
	for _, word := range words {
		wordFrequency[word]++
	}

	for key, val := range wordFrequency {
		fmt.Printf("%v\t\t%v\n", key, val)
	}
}

func main() {
	frequencyCheck()
}
