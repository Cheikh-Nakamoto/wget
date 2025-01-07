package utils

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func ConvertSize(size int64) string {
	var unit string
	var value float64

	switch {
	case size < 1024:
		unit = "bytes"
		value = float64(size)

	case size < 1024*1024:
		unit = "Kb"
		value = float64(size) / 1024

	case size < 1024*1024*1024:
		unit = "Mb"
		value = float64(size) / (1024 * 1024)

	case size < 1024*1024*1024*1024:
		unit = "Gb"
		value = float64(size) / (1024 * 1024 * 1024)

	default:
		unit = "Tb"
		value = float64(size) / (1024 * 1024 * 1024 * 1024)
	}

	return fmt.Sprintf("%.2f %s", value, unit)
}

func ConvertSpeed(speed float64) string {
	var value float64

	if speed < 1024 {
		return fmt.Sprintf("%.2f Kb/s", speed)
	} else {
		value = speed / 1024
		return fmt.Sprintf("%.2f Mb/s", value)
	}
}

func ToMega(size string) string {
	nb, _ := strconv.Atoi(size)
	ratio := math.Pow(10, float64(2))
	return fmt.Sprintf("%.2fMB", (math.Round(float64(nb)/(1024*1024)*ratio))/ratio)
}

func ParseRateLimit(input string) (int, error) {
	var multiplier int
	var value int

	switch {
	case input[len(input)-1] == 'k':
		multiplier = 1024
	case input[len(input)-1] == 'M':
		multiplier = 1024 * 1024
	default:
		return 0, fmt.Errorf("invalid rate limit format: %s", input)
	}

	fmt.Sscanf(input[:len(input)-1], "%d", &value)

	return value * multiplier, nil
}
