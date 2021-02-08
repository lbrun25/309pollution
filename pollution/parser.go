package pollution

import (
	"309pollution/utils"
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrTooManyArgs = errors.New("there are too many arguments")
	ErrNotEnoughArgs = errors.New("there are not enough arguments")
	ErrEmptyFile = errors.New("the file is empty")
	ErrSyntax = errors.New("text contains syntax error")
	ErrInvalidP = errors.New("'p' must be positive")
	ErrInvalidN = errors.New("'n' must be a positive integer and greater than 2")
	ErrInvalidX = errors.New("'x' must be a decimal greater or equal than 0 and inferior or equal to n - 1")
	ErrInvalidY = errors.New("'y' must be a decimal greater or equal than 0 and inferior or equal to n - 1")
)

const (
	maxArg = 4
	minArg = 4
)

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if arg == "-h" {
			return true
		}
	}
	return false
}

func retrievePoints(lines []string) ([]Point, error) {
	var points []Point

	for _, line := range lines {
		values := strings.Split(line, ";")

		if len(values) != 3 {
			return nil, ErrSyntax
		}

		res := map[int]int{}
		for i, value := range values {
			integer, err := strconv.Atoi(value); if err != nil {
				return nil, err
			}
			res[i] = integer
		}
		if res[2] < 0 {
			return nil, ErrInvalidP
		}
		points = append(points, Point{
			X: int64(res[0]),
			Y:  int64(res[1]),
			P:  int64(res[2]),
		})
	}

	return points, nil
}

// CheckArgs - check the user's args
func CheckArgs() (err error) {
	argsWithoutProg := os.Args[1:]

	// Check the number of arguments
	if len(argsWithoutProg) < minArg {
		return ErrNotEnoughArgs
	}
	if len(argsWithoutProg) > maxArg {
		return ErrTooManyArgs
	}

	// Retrieve args
	n, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil || n < 3 {
		return ErrInvalidN
	}
	x, err := strconv.ParseFloat(argsWithoutProg[2], 32); if err != nil {
		return ErrInvalidX
	}
	y, err := strconv.ParseFloat(argsWithoutProg[3], 32); if err != nil {
		return ErrInvalidY
	}
	N = int64(n)
	X = x
	Y = y

	// Get the content of the file
	lines, err := utils.ReadFile(argsWithoutProg[1]); if err != nil {
		return err
	}

	// Check if the file is empty
	if len(lines) == 0 {
		return ErrEmptyFile
	}

	// Retrieve values from the csv file
	points, err := retrievePoints(lines); if err != nil {
		return err
	}
	Points = points

	return nil
}