package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	FirstName string
	LastName  string
	Name      string
	NumSub    int
)

var Grades = make(map[string]float64)

func GradeCalculator() {
	fmt.Println("Welcome! Please enter your name. (e.g., Naol Daba)")
	fmt.Scanf("%s %s\n", &FirstName, &LastName)
	Name = FirstName + " " + LastName

	fmt.Printf("Hey %v, how many subjects did you take?\n", Name)
	fmt.Scanln(&NumSub)

	if NumSub <= 0 {
		fmt.Println("You did not take any subjects.")
		return
	}

	fmt.Printf("You took %d subjects.\n", NumSub)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < NumSub; i++ {
		fmt.Printf("Enter the name of subject number %v:\n", i+1)
		subject, _ := reader.ReadString('\n')
		subject = strings.TrimSpace(subject)
		var grade float64

		for {
			fmt.Printf("Enter the grade for %v (0-100):\n", subject)
			fmt.Scanln(&grade)
			if grade > 100 || grade < 0{
				fmt.Println("Invalid Input")
				fmt.Println()
			} else{
				break
			}
		}

		Grades[subject] = grade
	}

	fmt.Printf("Name\t\t%v\n", Name)

	var total float64

	for key, val := range Grades {
		fmt.Printf("%v\t\t%v\n", key, val)
		total += val
	}

	average := calcAverage(total, NumSub)

	fmt.Printf("Total\t\t%v\n", total)
	fmt.Printf("Average\t\t%v\n", average)
}

func calcAverage(total float64, numSub int) float64 {
	if numSub == 0 {
		return 0
	}
	return total / float64(numSub)
}

func main() {
	GradeCalculator()
}