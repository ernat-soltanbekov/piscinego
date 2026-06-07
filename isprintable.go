package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		// Если символ выходит за рамки диапазона [32, 126],
		// значит, он непечатный.
		if r < 32 || r > 126 {
			return false
		}
	}
	// Если прошли всю строку и не нашли "невидимок" — всё отлично
	return true
}
