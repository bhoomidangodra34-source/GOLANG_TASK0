package main

import (
	"fmt"
	"time"
)

func main() {

	// Q5.1
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	time.Sleep(time.Second)

	// Q5.2
	ch := make(chan string)

	go func() {
		ch <- "Hello from channel"
	}()

	message := <-ch
	fmt.Println(message)

	// Q5.3
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Q5.4
	defer fmt.Println("Deferred statement")

	fmt.Println("Main function is running")

	// Boss
	go func() {
		fmt.Println("Goroutine running")
	}()

	time.Sleep(time.Second)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}
