package calc

import (
	"fmt"
	"strconv"
)

const START string = "S"

// 0 - конец без ошибки
// 1 - закинуть в стек операций
// 2 - гененировать команду и читать следующий символ
// 3 - удалить из стека операций
// 4 - генерировать команду и повторить с тем же входным символом
// 5 - ошибка

//   + * ( ) " "
//   1 1 1 5  0
// + 2 1 1 4  4
// * 4 2 1 4  4
// ( 1 1 1 3  5

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func containsOperator(operators [8]string, target string) bool {
	for _, operator := range operators {
		if operator == target {
			return true
		}
	}
	return false
}

func execute(stack *[]float64, tokens *[]string, token string) (float64, error) {
	var operand1, operand2, result float64

	size := len(*stack)
	if len((*stack)) > 1 {
		operand1 = ((*stack)[size-2])
		operand2 = ((*stack)[size-1])
	} else {
		return 0, fmt.Errorf("недостаточно чисел для совершения операции")
	}
	if len(*stack) != 2 {
		*stack = (*stack)[:len(*stack)-2]
	} else {
		*stack = []float64{}
	}
	switch token {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result = operand1 / operand2
	}
	return result, nil
}

func doOne(tokens *[]string, token string) {
	// I – заслать символ в стек операций и читать следующий символ; ;
	*tokens = append(*tokens, token)
}

func doTwo(stack *[]float64, tokens *[]string, token string) error {
	// II – генерировать Kn, заслать в стек операций и читать следующий символ;
	previousToken := (*tokens)[len(*tokens)-1]
	*tokens = (*tokens)[:len(*tokens)-1]
	*tokens = append(*tokens, token)
	result, err := execute(stack, tokens, previousToken)
	if err != nil {
		return err
	}
	*stack = append(*stack, result)
	return nil
}

func doThree(tokens *[]string) {
	// III – удалить верхний символ из стека операций и читать следующий символ
	*tokens = (*tokens)[:len(*tokens)-1]
}

func doFour(stack *[]float64, tokens *[]string, token string) error {
	// IV – генерировать Кn и повторить с тем же входным символом
	previousToken := (*tokens)[len(*tokens)-1]
	result, err := execute(stack, tokens, previousToken)
	if err != nil {
		return err
	}
	*stack = append(*stack, result)
	*tokens = (*tokens)[:len(*tokens)-1]
	return nil
}

func evaluateRPN(expression string) (float64, error) {
	var float float64
	var numStr string
	var err error

	operators := [8]string{"S", "+", "-", "*", "/", "(", ")", "E"}

	operationsTable := map[string][7]int{
		//   +  -  *  /  (  )  E
		"S": {1, 1, 1, 1, 1, 5, 0}, // S
		"+": {2, 2, 1, 1, 1, 4, 4}, // +
		"-": {2, 2, 1, 1, 1, 4, 4}, // -
		"*": {4, 4, 2, 2, 1, 4, 4}, // *
		"/": {4, 4, 2, 2, 1, 4, 4}, // /
		"(": {1, 1, 1, 1, 1, 3, 5}, // (
	}

	stack := []float64{} // Числа
	tokens := []string{} // Операторы

	tokens = append(tokens, START)

	for _, char := range expression {
		token := string(char)

		// Считываем построчно цифры, чтобы получить число
		if isNumeric(token) {
			numStr += token
		} else if containsOperator(operators, token) {
			if numStr != "" { // Если попался оператор, то число закончилось, переводим во Float64
				float, _ = strconv.ParseFloat(numStr, 64)
				stack = append(stack, float)
				numStr = ""
			}
			for len(tokens) > 0 {
				previousToken := tokens[len(tokens)-1] // Смотрим относительно предыдущего символа
				if operationsTable[previousToken][operationsIndex(token)] == 1 {
					doOne(&tokens, token)
					break
				} else if operationsTable[previousToken][operationsIndex(token)] == 2 {
					err = doTwo(&stack, &tokens, token)
					break
				} else if operationsTable[previousToken][operationsIndex(token)] == 3 {
					doThree(&tokens)
					break
				} else if operationsTable[previousToken][operationsIndex(token)] == 4 {
					err = doFour(&stack, &tokens, token)
				} else {
					return 0, fmt.Errorf("ошибка в выражении")
				}
				if err != nil {
					return 0, err
				}
			}
		} else {
			return 0, fmt.Errorf("неверный символ: %c", char)
		}
	}
	if numStr != "" {
		float, _ = strconv.ParseFloat(numStr, 64)
		stack = append(stack, float)
		numStr = ""
	}
	for len(stack) != 1 && tokens[len(tokens)-1] != "S" {
		token := tokens[len(tokens)-1]
		err := doFour(&stack, &tokens, token)
		if err != nil {
			return 0, err
		}
	}

	return stack[0], nil
}

func operationsIndex(token string) int {
	operators := [7]string{"+", "-", "*", "/", "(", ")", "E"}
	for i, op := range operators {
		if op == token {
			return i
		}
	}
	return -1
}
