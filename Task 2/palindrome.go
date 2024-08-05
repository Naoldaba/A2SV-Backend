package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Palindrome() {
	fmt.Println("Welcome! Pls enter any string.")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	input := strings.ToLower(str)

	var cleanedRune []rune

	for _, char := range input {
		if unicode.IsDigit(char) || unicode.IsLetter(char){
			cleanedRune = append(cleanedRune, char)
		} 
	}

	n := len(cleanedRune)

	for i:=0; i<n/2; i++ {
		if cleanedRune[i] != cleanedRune[n-i-1]{
			fmt.Println("Not Palindrome")
			return
		}
	}
	fmt.Println("It is Palindrome")
}