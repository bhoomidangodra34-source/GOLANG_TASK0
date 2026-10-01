// student project CLI
package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	RollNo int
}

var students []Student
var studentMap = make(map[int]Student)

func addStudent() {
	var student Student

	fmt.Print("Enter name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Enter age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Enter roll number: ")
	fmt.Scan(&student.RollNo)

	students = append(students, student)
	studentMap[student.RollNo] = student

	fmt.Println("Student added")
}

func viewStudents() {
	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

	for _, student := range students {
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("Roll No:", student.RollNo)
		fmt.Println()
	}
}

func searchStudent() {
	var rollNo int

	fmt.Print("Enter roll number: ")
	fmt.Scan(&rollNo)

	student, exists := studentMap[rollNo]

	if exists {
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("Roll No:", student.RollNo)
	} else {
		fmt.Println("Student not found.")
	}
}
func main() {
	var choice int
	for {
		fmt.Println("\n--- Student Management CLI ---")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Search Student")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			viewStudents()
		case 3:
			searchStudent()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
