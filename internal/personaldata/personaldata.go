package personaldata

import "fmt"

// Параметры персоны
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Отображение данных персоны
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
