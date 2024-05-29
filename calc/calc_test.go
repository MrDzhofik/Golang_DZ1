package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Normal situations
func TestExpressions(t *testing.T) {
	expression := "(1+2)-3"
	expected := float64(0)

	got, _ := evaluateRPN(expression)

	// Сравниваем вывод с ожидаемым результатом

	assert.Equal(t, got, expected, "Простые выражения со скобками считаются верно")
}

func TestBrackets(t *testing.T) {
	expression := "4*(8/(9-1)-4)"
	expected := float64(-12)

	got, _ := evaluateRPN(expression)

	// Сравниваем вывод с ожидаемым результатом

	assert.Equal(t, got, expected, "Сложные выражения со скобками считаются верно")
}

func TestDigits(t *testing.T) {
	expression := "9+18/9*44-123"
	expected := float64(-26)

	got, _ := evaluateRPN(expression)

	// Сравниваем вывод с ожидаемым результатом

	assert.Equal(t, got, expected, "Числа считываются верно")
}

func TestZeroDivision(t *testing.T) {
	expression := "1+2/0"
	expected := "деление на ноль"

	_, err := evaluateRPN(expression)

	// Сравниваем вывод с ожидаемым результатом

	assert.Errorf(t, err, expected, "Деление на ноль распознается")
}

func TestWrongSymbol(t *testing.T) {
	expression := "1+!2/0"
	expected := "неверный символ"

	_, err := evaluateRPN(expression)

	// Сравниваем вывод с ожидаемым результатом

	assert.Errorf(t, err, expected, "Неверные выражения распознаются")
}

func TestNotEnoughDigit(t *testing.T) {
	expression := "(1-)"
	expected := "недостаточно чисел для совершения операции"

	_, err := evaluateRPN(expression)

	// Сравниваем вывод с ожидаемым результатом

	assert.Errorf(t, err, expected, "Неверные выражения распознаются")
}

func TestWrongExpression(t *testing.T) {
	expression := "1+2*6-99)(7+12-9*1284(9*(1+2)))"
	expected := "ошибка в выражении"

	_, err := evaluateRPN(expression)

	// Сравниваем вывод с ожидаемым результатом

	assert.Errorf(t, err, expected, "Неверные выражения распознаются")
}
