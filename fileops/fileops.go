package fileops

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("failed to find file")
	}

	fileContentString := string(file)
	valueFloat64, err := strconv.ParseFloat(fileContentString, 64)

	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return valueFloat64, nil
}
