package main
import "fmt"
func main() {
	var name string
	var subjectCount int

	fmt.Println("Enter student name:")
	fmt.Scanln(&name)

	fmt.Println("Enter number of subjects:")
	fmt.Scanln(&subjectCount)

	if subjectCount <= 0 {
		fmt.Println("Invalid number of subjects. Must be greater than zero.")
		return
	}

	subjects := make(map[string]float64)
	var total float64

	for i := 0; i < subjectCount; i++ {
		var subjectName string
		var grade float64

		fmt.Println("Enter name of subject:", i+1)
		fmt.Scanln(&subjectName)

		for {
			fmt.Println("Enter grade for", subjectName, "(0 - 100):")
			fmt.Scanln(&grade)

			if grade >= 0 && grade <= 100 {
				break
			} else {
				fmt.Println("Invalid grade. Please enter a value between 0 and 100.")
			}
		}
		subjects[subjectName] = grade
		total += grade
	}

	average := total / float64(subjectCount)

	fmt.Println("--------- Grade Report ----------")
	fmt.Println("Student Name:", name)

	for subject, grade := range subjects {
		fmt.Println("Subject:", subject, "- Grade:", grade)
	}

	fmt.Println("Average Grade:", average)
}
