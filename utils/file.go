package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// ReadFile - Read a file line by line.
// Based on laurent's answer https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
// It returns the lines of the file and any write error encountered.
func ReadFile(filepath string) ([]string, error) {
	var lines []string
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %s", filepath, "does not exist.\n"))
	}
	defer func() {
		err = file.Close()
	}()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("An error occurred when closing %s", filepath))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("An error occurred while reading %s", filepath))
	}
	return lines, nil
}