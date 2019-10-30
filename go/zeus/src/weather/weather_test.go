package weather

import (
	"reflect"
	"testing"
)

func TestHeadersParseFile(t *testing.T) {
	expected := []string{"14", "2"}
	minSpreads := MinimumSpread("../../weather.dat")
	if reflect.DeepEqual(minSpreads, expected) {
		t.Fatalf("Expected: %v\n Got: %v\n", expected, minSpreads)
	}
	t.Log("Pass", t.Name())
}
