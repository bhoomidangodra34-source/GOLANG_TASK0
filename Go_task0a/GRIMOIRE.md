 The gopher's Grimoire
## Realm 1 - Village of variables
### Q1.1    
The var keyword is used to declare a variaoble and it can be used both inside functions and also at package level so it can be accessed globally. Also in var keyword we need to explicitly give datatype while declaring .
The ':=' is used to declare a variable and initialize it directly. We don't need to explicitly give datatype while declaring Go automatically determine the datatype based on the value. also it can only be used inside a function.

### Q1.2
'const' is used to declare a variable whose value won't change in the whole code. if we want a variable to be the same we declare it using const variable. 
The main difference between 'const' and 'var' is that a variable declared using 'var' is it can be changed but if it is declare using 'const' can't be changed.

### Q1.3
int : it is basically used for whole numbers, positive numbers or negative numbers. 
float64 : it is used for decimal numbers for example 3.14,2.3,etc. 
string : it is used to store text or words or names. it is written in double quotes. 
bool : it is used to store boolean values such as true or false.
fmt.Println() : it is used to print something on screen. if we want to print some text it is written in double quotes and if we write variable name inside it will print its value.

### Q1.4
go automatically assigns a default value such as for int it is 0, for float64 it is 0.0 , for string it is empty string "" and for bool it is false.

### Realm Boss
Go has zero values because when we declare a variable without initializing it with a value then Go automatically assigns it a default zero value. when we want to use the variable later in the program it already has a predictable starting value. We can then use that variable without explicitly initializing it first, for example when using it in a loop where we want the initial value to start from zero.

### Realm 2 - The Forest of Control Flow
### Q2.1
If, else if and else in Go are basically the same as in other languages. Suppose there is some condition. First it will check the if condition. If the condition is satisfied, it will execute the code in the if block. Otherwise, it will go to the else if condition and check that. If that condition is also not satisfied, then it will go to the else block.

### Q2.2
Switch in Go is similar to switch in C. There are multiple cases, and based on the input, it will check the different cases. Like suppose we have case 1, case 2, case 3 and so on. If the input matches a case, it will execute the code inside that case. If none of the cases match, then it will execute the default block. Go also has a conditionless switch, where we don't give any value or condition after switch. Instead, we directly give conditions inside the cases, like if marks are greater than or equal to 90, then Grade A, if marks are greater than or equal to 75, then Grade B, and so on.

### Q2.3
There are three main ways of using the for loop in Go. The first one is the basic for loop, which is similar to the one we have in C. It has initialization, then a condition, and then increment or decrement. Suppose i is equal to 1, then we give the condition till where we want the loop to run, and then we increment i. The second one is the while-style for loop, where the initialization is outside and only the condition is given in the for loop, while the increment or decrement is done inside the code block. The third one is the infinite for loop, where we just write for and the loop keeps running until we use break to stop it.

### Q2.4
Range is used to go through the values of something like a slice or a map. Suppose a slice has five values, then range will go through all the five values. It can give us the index and the value of each element. In a map, range is used to access all the key-value pairs, so we can get the key as well as its corresponding value.

### Realm Boss
Go only has the for loop and it doesn't have separate while or do-while loops. I think the Go designers chose this to keep the language simple. Instead of having different loops for different purposes, we can use the same for loop in different ways, like a normal for loop, a while-style loop, or an infinite loop.

## Realm 3 - The castle of functions
### Q3.1
Named return values are basically when we make a function and we give a name to the return value. Like suppose we have sum = A + B, so instead of directly returning A plus B, we can give the return value a name like sum and then return it. So whenever we give a name to the return value in the function itself, it is called a named return value.

### Q3.2
Yes, a function can return more than one value in Go. Like in the example I have used, I have taken A and B as the parameters of type int, and I have returned (int, int). And inside the return, I have written A + B, A - B. So the first int will return the sum, that is A plus B, and the second int will return the difference, that is A minus B. So basically one function is returning two values.

### Q3.3
'...int' is called a variadic function. It is used when we want to give multiple parameters to a function. Like in my example, I have taken 'num ...int', so I can pass multiple integer values to the function. Then inside the function I have used a for loop to go through all the numbers and calculate their sum. Like if I give 10, 20 and 30, it will calculate 10 plus 20 plus 30.

### Q3.4
A closure is basically an anonymous function which can use the values from its surrounding scope. We don't need to give a name to the function. Like in my example, I have taken x = 10 outside the function and then inside I have taken y as a parameter and returned x + y. So even though I have not passed x as a parameter to the inner function, it can still use x because x is from the surrounding scope. So this is called a closure.

### Boss
In Go, result, error means that a function can return two things, one is the actual result and the other is the error. We can check the error using if (err != nil). If the error is not nil, it means that something went wrong, and then we can handle that error. I think Go uses this instead of exceptions because we can directly check and handle the error in the code instead of the error handling happening separately.

## Realm 4 - The dragon's lair of data structures 
### Q4.1
An array has a fixed size, whereas a slice has a dynamic size. So once we create an array with a particular size, that size cannot be changed. A slice can grow or shrink as needed. Both arrays and slices are mutable, which means we can change the values stored inside them.

### Q4.2
A map is basically a key-value pair. In `map[string]int`, the key will be of type string and the value will be of type int. Like suppose there is a map of student marks. The student's name can be the key, which is a string, and the marks can be the value, which is an int. So we can use the student's name to access their marks.

### Q4.3
A struct is like a structure which can have different data types inside it. Like suppose we have a student struct. It can have name as string, age as int and roll number as int. So these are three fields of the struct. Then we can create an instance of the student struct and give values to these fields, and then we can print that instance.

### Q4.4
A pointer receiver is basically like it points to the address of the original value. So the pointer receiver basically means that whatever changes we make through it, it will change the original value also.

### Boss
A value receiver is like it makes a copy of the original value, so if we make any changes, the changes only happen in that copy and not in the original one. But in a pointer receiver, it points to the original value, so if we make any changes, it will change in the original one also.

## Realm 5 - The Peak of Concurrency & Errors
### Q5.1
A goroutine basically allows a function to run concurrently with other work. We can create a goroutine by using the `go` keyword before the function call. So instead of waiting for the function to complete normally, Go can run it concurrently with other work.

### Q5.2
A channel is basically used to pass a value between two goroutines We first create a channel and then one goroutine can send a value using `ch <- value`. The other goroutine can receive that value using `<-ch`. 

### Q5.3
In golang the function also returns if their is some error along with the result . we can store them in variable like 'result ,err ' and then use 'if err != nil' followed by code to handle the error.

### Q5.4
defer is like do this when the current function is about to finish.

### Boss
Go uses goroutines and channels to make concurrent programming simpler and easier to manage. Goroutines are lightweight and allow multiple functions to run concurrently. Channels allow these goroutines to communicate and pass values between each other.

### FINAL BOSS - CLI
This program is a menu-driven Student Management CLI written in Go. The `Student` struct stores the student's name, age, and roll number. A slice stores all students, while a map uses the roll number to quickly search for a student. The `addStudent()` function adds student details, `viewStudents()` displays all students, and `searchStudent()` finds a student using their roll number. The `main()` function uses a loop and switch statement to display and handle the menu options.

