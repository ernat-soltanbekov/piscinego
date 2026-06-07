package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isBaseValid(base) {
		return
	}

	// Обработка отрицательных чисел
	if nbr < 0 {
		z01.PrintRune('-')
		// Используем uint, чтобы избежать переполнения при инверсии MinInt
		printRecursive(uint(nbr*-1), base)
	} else {
		printRecursive(uint(nbr), base)
	}
}

func printRecursive(n uint, base string) {
	b := uint(len(base))
	if n >= b {
		printRecursive(n/b, base)
	}
	z01.PrintRune(rune(base[n%b]))
}

func isBaseValid(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}
