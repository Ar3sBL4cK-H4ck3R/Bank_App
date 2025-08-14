package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatToFile(fileName string) (float64, error) {
	readFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Oops! Got some Error")
		return 0, errors.New("failed to find file")
	} else {
		floatValue, err := strconv.ParseFloat(string(readFile), 64)
		if err != nil {
			return 0, errors.New("float parsing error encountered")
		}
		return floatValue, nil
	}
}

func WirteFloatTofile(fileName string, valueString string) {
	os.WriteFile(fileName, []byte(valueString), 0644)
}
