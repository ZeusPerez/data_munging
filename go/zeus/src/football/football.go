package football

import (
	"strconv"
	"utils"
)

func MinimumGoals(path string) minGoals {
	footballData, _ := utils.ParseFile(path)
	initialTeam := footballData[1][1]
	initialGoals, _ := strconv.Atoi(footballData[1][6])
	initialGoalsAgainst, _ := strconv.Atoi(footballData[1][8])
	result := minGoals{initialTeam, intAbs(initialGoals - initialGoalsAgainst)}
	for i := 1; i < len(footballData); i++ {
		max, _ := strconv.Atoi(footballData[i][6])
		min, _ := strconv.Atoi(footballData[i][8])
		if intAbs(max-min) < result.diff {
			result.team = footballData[i][1]
			result.diff = intAbs(max - min)
		}

	}
	return result
}

type minGoals struct {
	team string
	diff int
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
