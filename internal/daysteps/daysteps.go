package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	splitDatastring := strings.Split(datastring, ",")
	if len(splitDatastring) != 2 {
		return fmt.Errorf("Данные должны содержать 2 элемента")
	}
	steps, err := strconv.Atoi(splitDatastring[0])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании количества шагов: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше 0")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(splitDatastring[1])
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании времени: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("время должно быть больше 0")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	way := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	spentWalkingCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, way, spentWalkingCalories)
	return result, nil
}
