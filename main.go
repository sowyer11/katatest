package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func romanToArabic(num string) int {
	rmap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return rmap[num]
}

func arabicToRoman(n int) string {
	v := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	r := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string
	for i := 0; i < len(v); i++ {
		for n >= v[i] {
			result += r[i]
			n -= v[i]
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение...")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода", err)
		return
	}

	pn := strings.Fields(input)
	if len(pn) != 3 {
		fmt.Println("Неправильный формат")
		return
	}

	var state string
	var first int
	var second int

	if romanToArabic(pn[0]) > 0 {
		if romanToArabic(pn[2]) > 0 {
			first, second = romanToArabic(pn[0]), romanToArabic(pn[2])
			state = "roman"
		} else {
			fmt.Println("Разные виды чисел")
			return
		}
	} else {
		if romanToArabic(pn[2]) > 0 {
			fmt.Println("Разные виды чисел")
			return
		}
	}

	if state != "roman" {
		first, err = strconv.Atoi(pn[0])
		if err != nil {
			fmt.Println("Неверное первое число")
			return
		}

		second, err = strconv.Atoi(pn[2])
		if err != nil {
			fmt.Println("Неверное второе число")
			return
		} else {
			state = "arabic"
		}
	}

	if state == "arabic" {
		if first > 10 {
			fmt.Println("Числа должны быть меньше 10")
			return
		}
		if second > 10 {
			fmt.Println("Числа должны быть меньше 10")
			return
		}
		if first <= 0 {
			fmt.Println("Числа должны быть больше 0")
			return
		}
		if second <= 0 {
			fmt.Println("Числа должны быть больше 0")
			return
		}
	}

	var success int
	var failure string

	switch pn[1] {
	case "+":
		if state == "roman" {
			failure = arabicToRoman(first + second)
		} else {
			success = first + second
		}
	case "-":
		if state == "roman" {
			if first-second > 0 {
				failure = arabicToRoman(first - second)
			} else {
				fmt.Println("Отрицательное число")
				return
			}
		} else {
			success = first - second
		}
	case "*":
		if state == "roman" {
			failure = arabicToRoman(first * second)
		} else {
			success = first * second
		}
	case "/":
		if second == 0 {
			fmt.Println("Деление на ноль")
			return
		}
		if state == "roman" {
			failure = arabicToRoman(first / second)
		} else {
			success = first / second
		}
	default:
		fmt.Println("Неверный оператор")
		return
	}

	if state == "roman" {
		fmt.Println("Результат: ", failure)
	} else {
		fmt.Println("Результат: ", success)
	}
}