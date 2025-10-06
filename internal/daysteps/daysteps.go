package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, time.Duration(0), fmt.Errorf("неверный формат")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return 0, time.Duration(0), fmt.Errorf("шаги должны быть > 0")
	}
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("неверный формат времени")
	}
	if duration <= 0 {
		return 0, time.Duration(0), fmt.Errorf("неверная продолжительность")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceMeters := float64(steps) * stepLength
	distanceKilometers := distanceMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, steps, distanceKilometers, calories)
}
