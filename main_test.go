package main

import (
	"fmt"
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

func TestZeroValue(t *testing.T) {
	list := []CaloriesCalculator{
		Swimming{
			Training: Training{
				TrainingType: "Плавание",
				Action:       0,
				LenStep:      0,
				Duration:     0,
				Weight:       0,
			},
			LengthPool: 0,
			CountPool:  0,
		},
		Walking{
			Training: Training{
				TrainingType: "Ходьба",
				Action:       0,
				LenStep:      0,
				Duration:     0,
				Weight:       0,
			},
			Height: 0,
		},
		Running{
			Training{
				TrainingType: "Бег",
				Action:       0,
				LenStep:      0,
				Duration:     0,
				Weight:       0,
			},
		},
	}
	data := []resultData{
		{
			duration:  0,
			distance:  0,
			meanSpeed: 0,
			calories:  0,
		},
		{
			duration:  0,
			distance:  0,
			meanSpeed: 0,
			calories:  0,
		},
		{
			duration:  0,
			distance:  0,
			meanSpeed: 0,
			calories:  0,
		},
	}
	for index, item := range list {
		expected := data[index]
		if errorMessage := checkError(item, expected); errorMessage != "" {
			t.Errorf(errorMessage)
		}
	}
}

func round2Dig(value float64) float64 {
	return math.Round(value*100) / 100
}

func checkError(item CaloriesCalculator, expected resultData) string {
	info := item.TrainingInfo()
	if duration := info.Duration.Minutes(); duration != expected.duration {
		return fmt.Sprintf("TestTrainingInfo - duration \"%s\": got %#v want \"%v\"", info.TrainingType, duration, expected.duration)
	}
	if distance := round2Dig(info.Distance); distance != expected.distance {
		return fmt.Sprintf("TestTrainingInfo - distance \"%s\": got %#v want \"%v\"", info.TrainingType, distance, expected.distance)
	}
	if meanSpeed := round2Dig(info.Speed); meanSpeed != expected.meanSpeed {
		return fmt.Sprintf("TestTrainingInfo - meanSpeed \"%s\": got %#v want \"%v\"", info.TrainingType, meanSpeed, expected.meanSpeed)
	}
	if calories := round2Dig(item.Calories()); calories != expected.calories {
		return fmt.Sprintf("TestTrainingInfo - calories \"%s\": got %#v want \"%v\"", info.TrainingType, calories, expected.calories)
	}
	return ""
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
		expected := data[index]
		if errorMessage := checkError(item, expected); errorMessage != "" {
			t.Errorf(errorMessage)
		}
	}
}
