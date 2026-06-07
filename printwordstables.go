package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, word := range table {
		if word != "" {
			// Выводим каждое слово посимвольно
			for _, char := range word {
				z01.PrintRune(char)
			}
			// После слова выводим перенос строки
			z01.PrintRune('\n')
		}
	}
}
