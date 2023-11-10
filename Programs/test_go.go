package main

import (
	"fmt"
)

func main() {
	fmt.Println("Калькулятор")
	calc()
}

func calc() {
	for {
		fmt.Println("Введите действие (+, -, *, /, закрыть )")
		var operation string
		fmt.Scan(&operation)
		if operation == "закрыть"{
			return
		}

		fmt.Println("Введите первое число")
		var firstInt int
		fmt.Scan(&firstInt)

		fmt.Println("Введите второе число число")
		var secondInt int
		fmt.Scan(&secondInt)

		fmt.Println("Ответ:")
		switch operation {
		case "+":
			fmt.Println(add(firstInt, secondInt))
		case "-":
			fmt.Println(noAdd(firstInt, secondInt))
		case "*":
			fmt.Println(mult(firstInt, secondInt))
		case "/":
			fmt.Println(div(firstInt, secondInt))
		}
	}
}

func add(x, y int) int {
	return x + y
}

func noAdd(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}
