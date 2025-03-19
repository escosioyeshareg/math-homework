package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) float64 {
	if b == 0 {
		return math.NaN()
	}
	return float64(a) / float64(b)
}

func main() {
	fmt.Println("Welcome to the Math Homework App!")

	var choice int
	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Add two numbers")
		fmt.Println("2. Subtract two numbers")
		fmt.Println("3. Multiply two numbers")
		fmt.Println("4. Divide two numbers")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var a, b int
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)
			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)
			result := add(a, b)
			fmt.Println("The sum of", a, "and", b, "is", result)
		case 2:
			var a, b int
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)
			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)
			result := subtract(a, b)
			fmt.Println("The difference between", a, "and", b, "is", result)
		case 3:
			var a, b int
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)
			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)
			result := multiply(a, b)
			fmt.Println("The product of", a, "and", b, "is", result)
		case 4:
			var a, b int
			fmt.Println("Enter the first number: ")
			fmt.Scan(&a)
			fmt.Println("Enter the second number: ")
			fmt.Scan(&b)
			result := divide(a, b)
			if result == math.NaN() {
				fmt.Println("Cannot divide by zero!")
			} else {
				fmt.Println("The result of", a, "/", b, "is", result)
			}
		default:
			fmt.Println("Invalid choice!")
		}
	}
}