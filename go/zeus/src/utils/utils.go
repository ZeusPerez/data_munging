package utils

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseFile(path string) ([][]string, error) {
	parsedFile, _ := FileToAry(path)
	cleanedFile := CleanNonAlphaNumeric(parsedFile)
	return RemoveEmptyLines(cleanedFile), nil
}

func FileToAry(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		lines = append(lines, fields)
	}
	return lines, scanner.Err()
}

func CleanNonAlphaNumeric(input [][]string) [][]string {
	var result [][]string

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	for _, element := range input {
		var line []string
		for _, subElement := range element {
			processedString := reg.ReplaceAllString(subElement, "")
			line = append(line, processedString)
		}
		result = append(result, line)
	}
	return result
}

func RemoveEmptyLines(input [][]string) [][]string {
	var result [][]string
	for _, element := range input {
		if len(element) > 0 {
			result = append(result, element)
		}
	}
	return result
}
