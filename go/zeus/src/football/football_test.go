package football

import (
	"testing"
)

func TestMinimumGoals(t *testing.T) {
	// expected := []int{14, 2}
	// minSpreads := MinimumSpread("../../weather.dat")
	// if reflect.DeepEqual(minSpreads, expected) {
	// 	t.Fatalf("Expected: %v\n Got: %v\n", expected, minSpreads)
	// }
	//
	expectedTeam := "AstonVilla"
	expectedDiff := 1
	minGoals := MinimumGoals("../../football.dat")
	if expectedTeam != minGoals.team {
		t.Fatalf("Expected team: %v\n Got: %v\n", expectedTeam, minGoals.team)
	}
	if expectedDiff != minGoals.diff {
		t.Fatalf("Expected diff: %v\n Got: %v\n", expectedDiff, minGoals.diff)
	}
	t.Log("Pass", t.Name())
}
