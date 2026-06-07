package piscine

// SortWordArr сортирует срез строк в порядке возрастания ASCII
func SortWordArr(a []string) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Сравниваем соседние строки
			if a[j] > a[j+1] {
				// Если текущий элемент больше следующего, меняем их местами
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
