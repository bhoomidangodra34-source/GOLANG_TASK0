package main

import "fmt"

func main() {

	// Q2.1
	marks := 75
	fmt.Println("marks is ", marks)
	if marks >= 90 {
		fmt.Println("Grade A")
	} else if marks >= 60 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// Q2.2
	Studentmarks := 60
	fmt.Println("marks is ", Studentmarks)
	switch {
	case Studentmarks >= 90:
		fmt.Println("GRADE A")
	case Studentmarks >= 75:
		fmt.Println("GRADE B")
	case Studentmarks >= 45:
		fmt.Println("GRADE C")
	default:
		fmt.Println("fail..")
	}

	// Q2.3
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	j := 1
	for {
		fmt.Println(j)
		j++

		if j > 5 {
			break
		}
	}

	// Q2.4
	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	marksMap := map[string]int{
		"bhoomi":  90,
		"arya":    80,
		"madhura": 85,
	}

	for name, mark := range marksMap {
		fmt.Println(name, mark)
	}
}
