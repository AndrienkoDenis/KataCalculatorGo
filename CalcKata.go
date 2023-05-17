package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result int
	convert := map[string]int{
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
	fmt.Println("Калькулятор умеет складывать, вычитать, умножать и делить\n" +
		"Введите два числа(римские или арабские)\n")

	// Чтение строки с консоли
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	// Преобразование строки в массив
	mass := strings.Split(scanner.Text(), " ")
	// если больше или меньше одной операции, неизвестная математическая операция
	if len(mass) != 3 {
		fmt.Println("Ошибка! Формат математической операции не удовлетворяет заданию")
		return
	}
	contain := strings.Contains("+ - * /", mass[1])
	// проверка операнта
	if !contain {
		fmt.Println("Ошибка!", mass[1], "- неизвестная математическая операция")
		return
	}
	// обработка содержимого
	a, num1 := strconv.Atoi(mass[0])
	b, num2 := strconv.Atoi(mass[2])
	o := mass[1]
	if num1 == nil && num2 == nil {
	} else if num1 != nil && num2 != nil {
		a = convert[mass[0]]
		b = convert[mass[2]]
		if a == 0 || b == 0 {
			fmt.Printf("Ошибка! Введены не подходящие числа")
			return
		}
	} else {
		fmt.Println("Ошибка! Используются одновременно разные системы счисления.")
		return
	}
	// Вычисление ответа
	switch {
	case a > 10 || b > 10 || a < 1 || b < 1:
		fmt.Println("Ошибка! Введены не подходящие числа")
		return
	case o == "+":
		result = a + b
	case o == "-":
		result = a - b
	case o == "*":
		result = a * b
	case o == "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("Ошибка! Деление на 0")
			return
		}
	}
	if num1 == nil && num2 == nil {
		fmt.Println(result)
	} else {
		if result < 1 {
			fmt.Println("Ошибка! Введены не подходящие числа")
			return
		}
		c1 := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
		c2 := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		var r2 string
		for i, j := range c1 {
			if result >= j {
				for result >= j {
					result -= j
					r2 += c2[i]
				}
			}
		}
		fmt.Println(r2)
	}
}
