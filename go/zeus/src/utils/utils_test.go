package utils

import (
	"reflect"
	"testing"
)

func TestHeadersParseFile(t *testing.T) {
	expected := []string{"Dy", "MxT", ",MnT", "AvT", "HDDay", "AvDP", "1HrP", "TPcpn", "WxType", "PDir", "AvSp", "Dir", "MxS", "SkyC", "MxR", "MnR", "AvSLP"}
	parsedFile, _ := ParseFile("../../weather.dat")
	if reflect.DeepEqual(expected, parsedFile[0]) {
		t.Fatalf("Expected: %v\n Got: %v\n", expected, parsedFile[0])
	}
	t.Log("Pass", t.Name())
}

func TestNoEmptyLineParseFile(t *testing.T) {
	parsedFile, _ := ParseFile("../../weather.dat")
	if len(parsedFile[1]) <= 0 {
		t.Fatalf("Expected no empty line after header")
	}
	t.Log("Pass", t.Name())
}

func TestNoAlphaNumericParseFile(t *testing.T) {
	parsedFile, _ := ParseFile("../../weather.dat")
	expected := "97"
	convertedVAlue := parsedFile[26][1]
	if expected != convertedVAlue {
		t.Fatalf("Expected only AlphaNumeric characters")
	}

	t.Log("Pass", t.Name())
}
