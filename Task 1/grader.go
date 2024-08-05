package main

import "fmt"

var (
	FistName string
	LastName string
	Name	string
	Num_sub	int
)

var Grades = make(map[string]float64)

func GradeCalculator(){
	fmt.Println("Welcome! Pls enter your name. (e.g Naol Daba)")
	fmt.Scanf("%s %s\n", &FistName, &LastName)
	Name = FistName + " "+ LastName

	fmt.Printf("Hey %v, how many subjects did you take?\n", Name)
	fmt.Scanln(&Num_sub)
	fmt.Printf("You took %d subjects.\n", Num_sub)

	var subject string
	var grade float64

	for i:=0; i<Num_sub; i++ {
		fmt.Printf("Enter name of subject number %v :\n", i+1)
		fmt.Scanln(&subject)

		fmt.Printf("Enter subject grade %v : (1-100)\n", i+1)
		fmt.Scanln(&grade)
		Grades[subject] = grade
	}

	fmt.Printf("Name\t\t%v\n", Name)

	var total, average float64

	for key, val := range Grades{
		fmt.Printf("%v\t\t%v\n", key, val)
		total += val
	}

	average = total / float64(Num_sub)

	fmt.Printf("total\t\t%v\n", total)
	fmt.Printf("average\t\t%v\n", average)
}

func main() {
	GradeCalculator()
}