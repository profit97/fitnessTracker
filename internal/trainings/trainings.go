package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	splitDatastring := strings.Split(datastring, ",")
	if len(splitDatastring) != 3 {
		return fmt.Errorf("Данные должны содержать 3 элемента")
	}
	steps, err := strconv.Atoi(splitDatastring[0])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании количества шагов: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше 0")
	}
	t.Steps = steps
	trainingType := splitDatastring[1]
	t.TrainingType = trainingType
	duration, err := time.ParseDuration(splitDatastring[2])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании времени: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("время должно быть больше 0")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	way := spentenergy.Distance(t.Steps, t.Personal.Height)
	avarageSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {

		return "", fmt.Errorf("неизвестный тип тренировки")

	}
	if t.TrainingType == "Ходьба" {
		spentWalkingCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), way, avarageSpeed, spentWalkingCalories)
		return result, nil
	}
	if t.TrainingType == "Бег" {
		spentRunningCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), way, avarageSpeed, spentRunningCalories)
		return result, nil
	}
	return "", nil
}
