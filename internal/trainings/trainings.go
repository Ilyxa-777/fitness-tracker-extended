package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training представляет собой тренировку пользователя.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse преобразует строку в структуру Training.
// Принимает строку с данными о тренировке, возвращает ошибку, если данные некорректны.
func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("invalid data string")
	}
	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return errors.New("invalid steps")
	}
	t.TrainingType = data[1]

	// if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {
	// 	return errors.New("неизвестный тип тренировки")
	// }

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return errors.New("invalid duration")
	}
	if t.Steps <= 0 {
		return errors.New("steps is less than 0")
	}
	if t.Duration <= 0 {
		return errors.New("duration is less than 0")
	}

	return nil
}

// ActionInfo возвращает информацию о тренировке.
// Принимает структуру Training, возвращает строку с информацией о тренировке и ошибку, если данные некорректны.
func (t Training) ActionInfo() (string, error) {
	var callories float64
	var err error

	if t.TrainingType == "Ходьба" {
		callories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else if t.TrainingType == "Бег" {
		callories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		spentenergy.Distance(t.Steps, t.Height),
		spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration),
		callories), nil
}
