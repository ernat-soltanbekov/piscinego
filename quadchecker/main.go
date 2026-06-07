package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. Читаем всё содержимое из стандартного ввода (stdin)
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}

	input := string(inputBytes)

	// Проверяем граничные случаи: пустой ввод или отсутствие перевода строки в конце
	if len(input) == 0 || input[len(input)-1] != '\n' {
		fmt.Println("Not a quad function")
		return
	}

	// 2. Разбиваем строку на массив строк по символу новой строки
	lines := strings.Split(input, "\n")
	
	// Так как строка заканчивалась на '\n', последний элемент слайса будет пустым. Удаляем его.
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Определяем размеры фигуры
	y := len(lines)      // Количество строк (высота)
	x := len(lines[0])   // Длина первой строки (ширина)

	if x == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Валидация: проверяем, что все строки имеют одинаковую длину
	for _, line := range lines {
		if len(line) != x {
			fmt.Println("Not a quad function")
			return
		}
	}

	// 3. Список всех доступных вариантов в алфавитном порядке
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var matches []string

	// Проверяем каждый тип quad на совпадение с входными данными
	for _, quad := range quads {
		if generateQuad(quad, x, y) == input {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad, x, y))
		}
	}

	// 4. Вывод результатов
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}

// Функция для генерации эталонного прямоугольника заданного типа и размеров
func generateQuad(kind string, x, y int) string {
	var sb strings.Builder
	for r := 0; r < y; r++ {
		for c := 0; c < x; c++ {
			sb.WriteByte(getChar(kind, r, c, x, y))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// Вспомогательная функция, возвращающая конкретный символ для координат (r, c)
func getChar(kind string, r, c, x, y int) byte {
	switch kind {
	case "quadA":
		if (r == 0 && c == 0) || (r == 0 && c == x-1) || (r == y-1 && c == 0) || (r == y-1 && c == x-1) {
			return 'o'
		}
		if r == 0 || r == y-1 {
			return '-'
		}
		if c == 0 || c == x-1 {
			return '|'
		}
		return ' '

	case "quadB":
		if r == 0 && c == 0 { return '/' }
		if r == 0 && c == x-1 { return '\\' }
		if r == y-1 && c == 0 { return '\\' }
		if r == y-1 && c == x-1 { return '/' }
		if r == 0 || r == y-1 || c == 0 || c == x-1 { return '*' }
		return ' '

	case "quadC":
		if r == 0 && c == 0 { return 'A' }
		if r == 0 && c == x-1 { return 'A' }
		if r == y-1 && c == 0 { return 'C' }
		if r == y-1 && c == x-1 { return 'C' }
		if r == 0 || r == y-1 || c == 0 || c == x-1 { return 'B' }
		return ' '

	case "quadD":
		if r == 0 && c == 0 { return 'A' }
		if r == 0 && c == x-1 { return 'C' }
		if r == y-1 && c == 0 { return 'A' }
		if r == y-1 && c == x-1 { return 'C' }
		if r == 0 || r == y-1 || c == 0 || c == x-1 { return 'B' }
		return ' '

	case "quadE":
		if r == 0 && c == 0 { return 'A' }
		if r == 0 && c == x-1 { return 'C' }
		if r == y-1 && c == 0 { return 'C' }
		if r == y-1 && c == x-1 { return 'A' }
		if r == 0 || r == y-1 || c == 0 || c == x-1 { return 'B' }
		return ' '
	}
	return ' '
}
