package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var str string

var Map = make(map[string]int)

func frequencyCheck() {
	fmt.Println("Welcome! Pls enter any string.")
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')

	input := strings.ToLower(str)
	var CleanedInput strings.Builder
	
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char){
			CleanedInput.WriteRune(char)
		} else {
			CleanedInput.WriteRune(' ')
		}
	}

	Words := strings.Fields(CleanedInput.String())
	Map := make(map[string]int)

    for _, word := range Words {
        Map[word]++
    }

    for key, val := range Map{
		fmt.Printf("%v\t\t%v\n",key, val)
	}
}


