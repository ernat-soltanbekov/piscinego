package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// 1. Случай для нуля
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// 2. Создаем "счетчик" для каждой цифры от 0 до 9
	counts := [10]int{}

	// 3. Обрабатываем число
	// Если число отрицательное, делаем его положительным
	if n < 0 {
		n = -n
	}

	for n > 0 {
		digit := n % 10 // получаем последнюю цифру
		counts[digit]++ // увеличиваем счетчик этой цифры
		n /= 10         // убираем последнюю цифру
	}

	// 4. Выводим цифры в порядке возрастания (от 0 до 9)
	for digit, count := range counts {
		for i := 0; i < count; i++ {
			// Превращаем индекс (цифру) в руну и печатаем
			z01.PrintRune(rune(digit + '0'))
		}
	}
}
