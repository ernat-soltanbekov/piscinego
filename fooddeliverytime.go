package piscine

// Определяем структуру, как требует задание
type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	// 1. Создаем "меню" через структуру
	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}

	// 2. Проверяем, какой заказ поступил
	switch order {
	case "burger":
		return burger.preptime
	case "chips":
		return chips.preptime
	case "nuggets":
		return nuggets.preptime
	default:
		// 3. Возвращаем 404, если блюда нет в меню
		return 404
	}
}
