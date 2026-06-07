package piscine

func ConcatParams(args []string) string {
	res := ""
	for i, arg := range args {
		res += arg
		// Добавляем \n только если это не последний элемент
		if i < len(args)-1 {
			res += "\n"
		}
	}
	return res
}
