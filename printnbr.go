package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Специальный случай для самого маленького возможного int
	if n == -9223372036854775808 {
		printStr("-9223372036854775808")
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

// Вспомогательная функция для вывода строки (для MinInt)
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
