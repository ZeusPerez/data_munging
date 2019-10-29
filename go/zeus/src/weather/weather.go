package weather

import (
	"strconv"
	"utils"
)

func MinimumSpread(path string) []int {
	weatherData, _ := utils.ParseFile(path)
	initialDay, _ := strconv.Atoi(weatherData[0][0])
	initialMax, _ := strconv.Atoi(weatherData[0][1])
	initialMin, _ := strconv.Atoi(weatherData[0][2])
	result := []int{initialDay, initialMax - initialMin}
	for _, element := range weatherData[1:] {
		max, _ := strconv.Atoi(element[1])
		min, _ := strconv.Atoi(element[2])
		if (max - min) < result[1] {
			day, _ := strconv.Atoi(element[0])
			result = []int{day, max - min}
		}

	}
	return result
}
