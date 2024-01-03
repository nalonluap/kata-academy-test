package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BinaryCalc struct {
	A        int
	B        int
	Roman    bool
	Operator string
}

func (calc *BinaryCalc) Input() {
	var a, b, op string
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	splitLine := strings.Split(line, " ")
	if err != nil {
		fmt.Print("Ошибка ввода через консоль\n")
	}

	if len(splitLine) == 1 {
		fmt.Print("Строка не является математической операцией\n")
		return
	} else if len(splitLine) > 3 {
		fmt.Print("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)\n")
		return
	}

	a, op, b = splitLine[0], splitLine[1], strings.TrimSuffix(splitLine[2], "\n")

	var x, y = convertToArabic(a), convertToArabic(b)

	if x == 0 && y == 0 {
		calc.A = convertToNumber(a)
		calc.B = convertToNumber(b)
		calc.Operator = op
		return
	}

	if x == 0 || y == 0 {
		fmt.Print("Нельзя использовать одновременно разные системы счисления\n")
		return
	}

	calc.A = x
	calc.B = y
	calc.Roman = true
	calc.Operator = op
}

func (calc *BinaryCalc) Output(result int) {
	if calc.Roman {
		if result < 1 {
			fmt.Println("В римской системе нет отрицательных чисел")
			return
		}
		fmt.Println(convertToRoman(result))
		return
	}
	fmt.Println(result)
}

func (calc *BinaryCalc) Calculator() {
	switch calc.Operator {
	case "+":
		calc.Output(calc.Add())
	case "-":
		calc.Output(calc.Sub())
	case "*":
		calc.Output(calc.Multi())
	case "/":
		calc.Output(calc.Div())
	}
}

func (calc *BinaryCalc) Add() int {
	return calc.A + calc.B
}

func (calc *BinaryCalc) Sub() int {
	return calc.A - calc.B
}

func (calc *BinaryCalc) Multi() int {
	return calc.A * calc.B
}

func (calc *BinaryCalc) Div() int {
	return calc.A / calc.B
}

func convertToArabic(number string) int {
	var roman = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	return roman[number]
}

func convertToRoman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func convertToNumber(n string) int {
	var number int
	var err error
	number, err = strconv.Atoi(n)
	if err != nil {
		fmt.Println("Введен некорректный формат числа", err, "\n")
	}

	return number
}

func main() {
	calc := BinaryCalc{}

	calc.Input()
	calc.Calculator()
}
