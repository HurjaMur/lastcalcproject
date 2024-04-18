package agent

import (
	"strconv"
	"strings"
)

// ProcessExpression вычисляет математическое выражение и возвращает результат
func ProcessExpression(expression string) int {
	// Выделение чисел и операций из выражения
	tokens := strings.FieldsFunc(expression, func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	})

	nums := make([]int, 0)
	operations := make([]string, 0)

	for _, token := range tokens {
		num, _ := strconv.Atoi(token)
		nums = append(nums, num)
	}

	for _, c := range expression {
		char := string(c)
		if char == "+" || char == "-" || char == "*" || char == "/" {
			operations = append(operations, char)
		}
	}

	// Обработка операций с учетом приоритета
	// Сначала умножение и деление
	for i := 0; i < len(operations); i++ {
		if operations[i] == "*" {
			nums[i] *= nums[i+1]
			nums = append(nums[:i+1], nums[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--
		} else if operations[i] == "/" {
			nums[i] /= nums[i+1]
			nums = append(nums[:i+1], nums[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--
		}
	}

	// Затем сложение и вычитание
	result := nums[0]

	for i, op := range operations {
		switch op {
		case "+":
			result += nums[i+1]
		case "-":
			result -= nums[i+1]
		}
	}

	return result
}
