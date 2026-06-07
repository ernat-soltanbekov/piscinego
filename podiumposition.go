package piscine

func PodiumPosition(podium [][]string) [][]string {
	// Сохраняем длину среза в переменную
	n := len(podium)

	// Нам нужно дойти только до середины массива (n/2)
	for i := 0; i < n/2; i++ {
		// Меняем местами симметричные элементы (первый с последним, второй с предпоследним)
		podium[i], podium[n-1-i] = podium[n-1-i], podium[i]
	}

	// Возвращаем измененный срез
	return podium
}
