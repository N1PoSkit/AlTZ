package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите выражение (например, \"hello\" + \"world\"): ")
	fmt.Scanln(&input)

	// Удаляем пробелы
	input = strings.ReplaceAll(input, " ", "")

	// Определяем операцию
	if strings.Contains(input, "+") {
		handleAddition(input)
	} else if strings.Contains(input, "-") {
		handleSubtraction(input)
	} else if strings.Contains(input, "*") {
		handleMultiplication(input)
	} else if strings.Contains(input, "/") {
		handleDivision(input)
	} else {
		panic("Неподдерживаемая операция")
	}
}

func handleAddition(input string) {
	parts := strings.Split(input, "+")
	if len(parts) != 2 {
		panic("Неправильный формат сложения")
	}
	str1, str2 := extractStrings(parts[0], parts[1])
	result := str1 + str2
	printResult(result)
}

func handleSubtraction(input string) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		panic("Неправильный формат вычитания")
	}
	str1, str2 := extractStrings(parts[0], parts[1])
	result := strings.ReplaceAll(str1, str2, "")
	printResult(result)
}

func handleMultiplication(input string) {
	parts := strings.Split(input, "*")
	if len(parts) != 2 {
		panic("Неправильный формат умножения")
	}
	str1 := extractString(parts[0])
	num := parseNumber(parts[1])
	result := strings.Repeat(str1, num)
	printResult(result)
}

func handleDivision(input string) {
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		panic("Неправильный формат деления")
	}
	str1 := extractString(parts[0])
	num := parseNumber(parts[1])
	if num <= 0 {
		panic("Деление на ноль недопустимо")
	}
	result := str1[:len(str1)/num]
	printResult(result)
}

func extractStrings(part1, part2 string) (string, string) {
	str1 := extractString(part1)
	str2 := extractString(part2)
	return str1, str2
}

func extractString(part string) string {
	if strings.HasPrefix(part, "\"") && strings.HasSuffix(part, "\"") {
		return part[1 : len(part)-1]
	}
	panic("Ожидалась строка в кавычках")
}

func parseNumber(part string) int {
	numStr := extractString(part)
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 1 || num > 10 {
		panic("Число должно быть от 1 до 10")
	}
	return num
}

func printResult(result string) {
	if len(result) > 40 {
		result = result[:40] + "..."
	}
	fmt.Println("Результат:", result)
}
