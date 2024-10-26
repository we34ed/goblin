package main

import (
	"fmt"
	"time"
)

// Общие константы для вычислений.
const (
	MInKm      = 1000 // количество метров в одном километре
	MinInHours = 60   // количество минут в одном часе
	LenStep    = 0.65 // длина одного шага
	CmInM      = 100  // количество сантиметров в одном метре
)

// Training общая структура для всех тренировок
type Training struct {
	trainingType string        // тип тренировки
	action       int           // количество повторов(шаги, гребки при плавании)
	lenStep      float64       // длина одного шага или гребка в м
	duration     time.Duration // продолжительность тренировки
	weight       float64       // вес пользователя в кг
}

// distance возвращает дистанцию, которую преодолел пользователь.
// Формула расчета:
// количество_повторов * длина_шага / м_в_км
func (t Training) distance() float64 {

	return float64(t.action) * LenStep / MInKm
}

// meanSpeed возвращает среднюю скорость бега или ходьбы.
func (t Training) meanSpeed() float64 {
	if t.duration == 0 {
		return 0
	}
	return t.distance() / t.duration.Hours()
}

// Calories возвращает количество потраченных килокалорий на тренировке.
// Пока возвращаем 0, так как этот метод будет переопределяться для каждого типа тренировки.
func (t Training) Calories() float64 {

	return 0
}

// InfoMessage содержит информацию о проведенной тренировке.
type InfoMessage struct {
	trainingType string        // тип тренировки
	duration     time.Duration // длительность тренировки
	distance     float64       // расстояние, которое преодолел пользователь
	speed        int           // средняя скорость, с которой двигался пользователь
	calories     float64       // количество потраченных килокалорий на тренировке
}

// TrainingInfo возвращает структуру InfoMessage, в которой хранится вся информация о проведенной тренировке.
func (t Training) TrainingInfo() InfoMessage {

	return InfoMessage{}
}

// String возвращает строку с информацией о проведенной тренировке.
func (i InfoMessage) String() string {
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.0f мин\nДистанция: %.2f км.\nСр. скорость: %d км/ч\nПотрачено ккал: %.2f\n",
		i.trainingType,
		i.duration.Minutes(),
		i.distance,
		i.speed,
		i.calories,
	)
}

// CaloriesCalculator интерфейс для структур: Running, Walking и Swimming.
type CaloriesCalculator interface {
	Calories() float64
	TrainingInfo() InfoMessage
}

// Константы для расчета потраченных килокалорий при беге.
const (
	CaloriesMeanSpeedMultiplier = 18   // множитель средней скорости бега
	CaloriesMeanSpeedShift      = 1.79 // коэффициент изменения средней скорости
)

// Running структура, описывающая тренировку Бег.
type Running struct {
	// добавьте необходимые поля в структуру
	Training
}

// Calories возввращает количество потраченных килокалория при беге.
// Формула расчета:
// ((18 * средняя_скорость_в_км/ч + 1.79) * вес_спортсмена_в_кг / м_в_км * время_тренировки_в_часах * мин_в_часе)
// Это переопределенный метод Calories() из Training.
func (r Running) Calories() float64 {
	speed := r.meanSpeed()
	calories := ((CaloriesMeanSpeedMultiplier*speed + CaloriesMeanSpeedShift) * r.weight / MInKm * r.duration.Hours() * MinInHours)
	return calories
}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (r Running) TrainingInfo() InfoMessage {
	return InfoMessage{
		trainingType: r.trainingType,
		duration:     r.duration,
		distance:     r.distance(),
		speed:        int(r.meanSpeed()),
		calories:     r.Calories(),
	}
}

// Константы для расчета потраченных килокалорий при ходьбе.
const (
	CaloriesWeightMultiplier      = 0.035 // коэффициент для веса
	CaloriesSpeedHeightMultiplier = 0.029 // коэффициент для роста
	KmHInMsec                     = 0.278 // коэффициент для перевода км/ч в м/с
)

// Walking структура описывающая тренировку Ходьба
type Walking struct {
	Training
	height float64
}

// Calories возвращает количество потраченных килокалорий при ходьбе.
// Формула расчета:
// ((0.035 * вес_спортсмена_в_кг + (средняя_скорость_в_метрах_в_секунду**2 / рост_в_метрах)
// * 0.029 * вес_спортсмена_в_кг) * время_тренировки_в_часах * мин_в_ч)
// Это переопределенный метод Calories() из Training.
func (w Walking) Calories() float64 {
	speed := w.meanSpeed() * KmHInMsec
	heightInM := float64(w.height) / CmInM
	return (CaloriesWeightMultiplier*w.weight + (speed*speed/heightInM)*CaloriesSpeedHeightMultiplier*w.weight) * w.duration.Hours()
}

// TrainingInfo возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (w Walking) TrainingInfo() InfoMessage {
	return InfoMessage{
		trainingType: w.trainingType,
		duration:     w.duration,
		distance:     w.distance(),
		speed:        int(w.meanSpeed()),
		calories:     w.Calories(),
	}
}

// Константы для расчета потраченных килокалорий при плавании.
const (
	SwimmingLenStep                  = 1.38 // длина одного гребка
	SwimmingCaloriesMeanSpeedShift   = 1.1  // коэффициент изменения средней скорости
	SwimmingCaloriesWeightMultiplier = 2    // множитель веса пользователя
)

// Swimming структура, описывающая тренировку Плавание
type Swimming struct {
	Training
	lengthPool int // длина бассейна
	countPool  int // количество пересечений бассейна
}

// meanSpeed возвращает среднюю скорость при плавании.
// Формула расчета:
// длина_бассейна * количество_пересечений / м_в_км / продолжительность_тренировки
// Это переопределенный метод Calories() из Training.
func (s Swimming) meanSpeed() float64 {
	distanceKm := float64(s.lengthPool*s.countPool) / MInKm

	speed := distanceKm / s.duration.Hours()

	return speed
}

// Calories возвращает количество калорий, потраченных при плавании.
// Формула расчета:
// (средняя_скорость_в_км/ч + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * вес_спортсмена_в_кг * время_тренировки_в_часах
// Это переопределенный метод Calories() из Training.
func (s Swimming) Calories() float64 {
	speed := s.meanSpeed()
	return (speed + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * s.weight * s.duration.Hours()
}

// TrainingInfo returns info about swimming training.
// Это переопределенный метод TrainingInfo() из Training.
func (s Swimming) TrainingInfo() InfoMessage {
	return InfoMessage{
		trainingType: s.trainingType,
		duration:     s.duration,
		distance:     s.distance(),
		speed:        int(s.meanSpeed()),
		calories:     s.Calories(),
	}
}

// ReadData возвращает информацию о проведенной тренировке.
func ReadData(training CaloriesCalculator) string {

	calories := training.Calories()

	info := training.TrainingInfo()

	info.calories = calories

	return fmt.Sprint(info)
}

func main() {
	swimming := Swimming{
		Training: Training{
			trainingType: "Плавание",
			action:       2000,
			lenStep:      SwimmingLenStep,
			duration:     90 * time.Minute,
			weight:       85,
		},
		lengthPool: 50,
		countPool:  5,
	}

	fmt.Println(ReadData(swimming))

	walking := Walking{
		Training: Training{
			trainingType: "Ходьба",
			action:       20000,
			lenStep:      LenStep,
			duration:     3*time.Hour + 45*time.Minute,
			weight:       85,
		},
		height: 185,
	}

	fmt.Println(ReadData(walking))

	running := Running{
		Training: Training{
			trainingType: "Бег",
			action:       5000,
			lenStep:      LenStep,
			duration:     30 * time.Minute,
			weight:       85,
		},
	}

	fmt.Println(ReadData(running))
	fmt.Println("привет Ваня")
	fmt.Println("111")
}
