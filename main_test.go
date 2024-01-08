package main

import (
	"math"
	"testing"
	"time"
)

type resultData struct {
	duration  float64
	distance  float64
	meanSpeed float64
	calories  float64
}

func TestTrainingInfo(t *testing.T) {
	list := []CaloriesCalculator{
		Swimming{
			Training: Training{
				TrainingType: "Плавание",
				Action:       2000,
				LenStep:      SwimmingLenStep,
				Duration:     90 * time.Minute,
				Weight:       85,
			},
			LengthPool: 50,
			CountPool:  5,
		},
		Walking{
			Training: Training{
				TrainingType: "Ходьба",
				Action:       20000,
				LenStep:      LenStep,
				Duration:     3*time.Hour + 45*time.Minute,
				Weight:       85,
			},
			Height: 185,
		},
		Running{
			Training: Training{
				TrainingType: "Бег",
				Action:       5000,
				LenStep:      LenStep,
				Duration:     30 * time.Minute,
				Weight:       85,
			},
		},
	}
	data := []resultData{
		{
			duration:  90,
			distance:  2.76,
			meanSpeed: 0.17,
			calories:  323,
		},
		{
			duration:  225,
			distance:  13,
			meanSpeed: 3.47,
			calories:  947.82,
		},
		{
			duration:  30,
			distance:  3.25,
			meanSpeed: 6.5,
			calories:  302.91,
		},
	}

	for index, item := range list {
		duration := item.TrainingInfo().Duration.Minutes()
		distance := math.Round(item.TrainingInfo().Distance*100) / 100
		meanSpeed := math.Round(item.TrainingInfo().Speed*100) / 100
		calories := math.Round(item.Calories()*100) / 100
		trainingType := item.TrainingInfo().TrainingType
		if duration != data[index].duration {
			t.Errorf("TestTrainingInfo - duration \"%s\": got %#v want \"%v\"", trainingType, duration, data[index].duration)
		} else if distance != data[index].distance {
			t.Errorf("TestTrainingInfo - distance \"%s\": got %#v want \"%v\"", trainingType, distance, data[index].distance)
		} else if meanSpeed != data[index].meanSpeed {
			t.Errorf("TestTrainingInfo - meanSpeed \"%s\": got %#v want \"%v\"", trainingType, meanSpeed, data[index].meanSpeed)
		} else if calories != data[index].calories {
			t.Errorf("TestTrainingInfo - calories \"%s\": got %#v want \"%v\"", trainingType, calories, data[index].calories)
		}
	}
}
