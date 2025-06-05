package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories вычисляет количество калорий, потраченных при ходьбе.
// Принимает количество шагов, вес и рост пользователя, а также продолжительность ходьбы.
// Возвращает количество потраченных калорий и ошибку, если входные параметры некорректны.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0, errors.New("duration is less than 0")
	}
	if steps <= 0 {
		return 0, errors.New("steps is less than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight is less than 0")
	}
	if height <= 0 {
		return 0, errors.New("height is less than 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	callories := ((weight * meanSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient
	return callories, nil
}

// RunningSpentCalories вычисляет количество калорий, потраченных при беге.
// Принимает количество шагов, вес и рост пользователя, а также продолжительность бега.
// Возвращает количество потраченных калорий и ошибку, если входные параметры некорректны.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0, errors.New("duration is less than 0")
	}
	if steps <= 0 {
		return 0, errors.New("steps is less than 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight is less than 0")
	}
	if height <= 0 {
		return 0, errors.New("height is less than 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	callories := (weight * meanSpeed * duration.Minutes()) / minInH
	return callories, nil
}

// MeanSpeed вычисляет среднюю скорость движения на основе количества шагов,
// роста пользователя и продолжительности активности.
// Возвращает скорость в километрах в час.
// Если продолжительность активности меньше или равна 0, возвращает 0.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	speed := distance / duration.Hours()
	return speed
}

// Distance вычисляет расстояние, пройденное пользователем на основе количества шагов и роста.
// Принимает количество шагов и рост пользователя, возвращает расстояние в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := stepLength * float64(steps)
	distanceKm := distance / mInKm

	return distanceKm
}
