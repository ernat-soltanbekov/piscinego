package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	// Если строки не равны и а не меньше b, значит а > b
	return 1
}
