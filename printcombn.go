package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	result := make([]rune, n)
	for i := 0; i < n; i++ {
		result[i] = '0' + rune(i)
	}

	first := true
	for {
		if !first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for _, r := range result {
			z01.PrintRune(r)
		}
		first = false

		i := n - 1
		// Приводим все части к типу rune для сравнения
		for i >= 0 && result[i] == '9'-rune(n-1-i) {
			i--
		}
		if i < 0 {
			break
		}
		result[i]++
		for j := i + 1; j < n; j++ {
			result[j] = result[j-1] + 1
		}
	}
	z01.PrintRune('\n')
}
