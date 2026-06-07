package piscine

// ForEach принимает функцию f и срез чисел a.
// Она поочередно применяет f к каждому элементу среза a.
func ForEach(f func(int), a []int) {
	for _, val := range a {
		f(val)
	}
}
