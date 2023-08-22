package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор")
	fmt.Println("Введите выражение:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()
	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

func calculate(expression string) (float64, error) {
	tokens := strings.Fields(expression)
	if len(tokens) != 3 {
		return 0, fmt.Errorf("некорректное выражение")
	}

	num1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, fmt.Errorf("некорректное число")
	}

	num2, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0, fmt.Errorf("некорректное число")
	}

	switch tokens[1] {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("некорректная операция")
	}
}
