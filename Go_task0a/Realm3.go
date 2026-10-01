package main

import "fmt"

//Q3.1
func sumofno(a int, b int) (sumofno int) {
	sumofno = a + b
	return
}

//Q3.2
func sumdiff(a int, b int) (int, int) {
	return a + b, a - b
}

//Q3.3
func sumofmultiple(num ...int) int {
	total := 0
	for _, number := range num {
		total += number
	}
	return total
}

func main() {
	//Q3.1
	fmt.Println("Sum of 10 and 20 is ", sumofno(10, 20))
	//Q3.2
	sum, diff := sumdiff(10, 5)
	fmt.Println("sum and difference is", sum, diff)
	//Q3.3
	fmt.Println("Sum of 10, 20 and 30 is ", sumofmultiple(10, 20, 30))
	//Q3.4
	x := 10
	add := func(y int) int {
		return x + y
	}
	fmt.Println(add(5))

}
