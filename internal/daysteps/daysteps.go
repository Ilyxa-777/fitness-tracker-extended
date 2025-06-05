package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse преобразует строку в структуру DaySteps.
// Принимает строку с данными о количестве шагов и продолжительности дня, возвращает ошибку, если данные некорректны.
func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("invalid data string")
	}
	ds.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return errors.New("invalid steps")
	}
	ds.Duration, err = time.ParseDuration(data[1])
	if err != nil {
		return errors.New("invalid duration")
	}
	if ds.Steps <= 0 {
		return errors.New("steps is less than 0")
	}
	if ds.Duration <= 0 {
		return errors.New("duration is less than 0")
	}

	return nil
}

// ActionInfo возвращает информацию о количестве шагов, дистанции и количестве сожженных калорий.
// Принимает структуру DaySteps, возвращает строку с информацией о количестве шагов, дистанции и количестве сожженных калорий и ошибку, если данные некорректны.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	callories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		callories), nil
}
