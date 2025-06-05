package actioninfo

import (
	"fmt"
)

type DataParser interface {
	// Parse преобразует строку в структуру данных
	Parse(datastring string) error

	// ActionInfo возвращает информацию о действии
	ActionInfo() (string, error)
}

// Info обрабатывает набор данных о действиях пользователя.
// Принимает слайс строк с данными и парсер данных.
// Для каждой строки выводит информацию об активности или логирует ошибку.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Парсим данные
		if err := dp.Parse(data); err != nil {
			fmt.Printf("Error parsing data: %v", err)
			continue
		}

		// Получаем и выводим информацию
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Error getting action info: %v", err)
			continue
		}

		fmt.Print(info)
	}
}
