package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	calc()
}

func calc() {
	for {
		fmt.Println("Введите первое число")
		var firstInt int
		fmt.Scan(&firstInt)
		fmt.Println(firstInt)
		fmt.Println("Введите второе число число")
		var secondInt int
		fmt.Scan(&secondInt)
		fmt.Println(secondInt)
		fmt.Println(add(firstInt, secondInt))
		return
	}
}

func add(x, y int) int {
	return x + y
}

func noAdd(x, y int) int {
	return x - y
}
