package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", time.Duration(0), fmt.Errorf("неверный формат")
	}
	step, err := strconv.Atoi(strings.TrimSpace(slice[0]))
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("неверный формат шагов")
	}
	if step <= 0 {
		return 0, "", time.Duration(0), fmt.Errorf("неверный формат шагов")
	}
	duration, err := time.ParseDuration(strings.TrimSpace(slice[2]))
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("неверный формат времени")
	}
	if duration <= 0 {
		return 0, "", time.Duration(0), fmt.Errorf("неверная продолжительность")
	}
	activity := strings.TrimSpace(slice[1])
	return step, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	stepsCompleted := float64(steps) * stepLength
	interval := stepsCompleted / mInKm
	return interval
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	interval := distance(steps, height)
	averageSpeed := interval / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	step, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	interval := distance(step, height)
	averageSpeed := meanSpeed(step, height, duration)
	switch activity {
	case "Бег":
		runCalories, err := RunningSpentCalories(step, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, activity, duration.Hours(), interval, averageSpeed, runCalories), nil
	case "Ходьба":
		walkCalories, err := WalkingSpentCalories(step, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, activity, duration.Hours(), interval, averageSpeed, walkCalories), nil
	}
	return "", fmt.Errorf("неизвестный тип тренировки")
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("неверное количество шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("неверный вес")
	}
	if height <= 0 {
		return 0, fmt.Errorf("неверный рост")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверное продолжительность бега")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * weight * averageSpeed) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функциюпщ
	if steps <= 0 {
		return 0, fmt.Errorf("неверное количество шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("неверный вес")
	}
	if height <= 0 {
		return 0, fmt.Errorf("неверный рост")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверное продолжительность бега")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * weight * averageSpeed) / minInH * walkingCaloriesCoefficient
	return calories, nil
}
