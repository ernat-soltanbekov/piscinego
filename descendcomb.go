package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	// Внешний цикл для первого числа (от 99 до 01)
	for i := 99; i >= 0; i-- {
		// Внутренний цикл для второго числа (всегда меньше первого)
		for j := i - 1; j >= 0; j-- {
			// Вывод первого числа
			z01.PrintRune(rune(i/10 + '0'))
			z01.PrintRune(rune(i%10 + '0'))

			// Пробел между числами в паре
			z01.PrintRune(' ')

			// Вывод второго числа
			z01.PrintRune(rune(j/10 + '0'))
			z01.PrintRune(rune(j%10 + '0'))

			// Проверка: если это не последняя пара (01 00), выводим запятую и пробел
			if !(i == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
