package main
import "fmt"
type Student struct {
	Name   string
	Age    int
	RollNo int
}
// Q4.4
func (s *Student) changeAge(age int) {
	s.Age = age
}
func main() {
	// Q4.1
	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println(numbers)
	values := []int{10, 20, 30}
	values = append(values, 40)
	fmt.Println(values)

	// Q4.2
	marks := map[string]int{
		"Bhoomi":  90,
		"arya":    80,
		"madhura": 85,
	}
	fmt.Println(marks)
	fmt.Println(marks["Bhoomi"])

	// Q4.3
	student1 := Student{
		Name:   "Bhoomi",
		Age:    18,
		RollNo: 033,
	}
	fmt.Println(student1)

	// Q4.4
	student1.changeAge(19)
	fmt.Println(student1)

	// Boss
	student1.changeAge(20)
	fmt.Println(student1)
}
