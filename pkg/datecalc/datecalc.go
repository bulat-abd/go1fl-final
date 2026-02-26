package datecalc

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"
)

func AfterNow(date, now time.Time) bool {
	// consider dates only!
	return date.Truncate(24 * time.Hour).After(now.Truncate(24 * time.Hour))
}

func processDailyRepeat(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	parts := strings.Split(repeat, " ")
	if len(parts) < 2 {
		return "", errors.New("Bad repeat parameter value")
	}
	days, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}
	if days > 400 {
		return "", errors.New("D value of repeat parameter is too high")
	}
	for {
		date = date.AddDate(0, 0, days)
		if AfterNow(date, now) {
			break
		}
	}
	return date.Format("20060102"), nil
}

func processYearlyRepeat(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	for {
		date = date.AddDate(1, 0, 0)
		if AfterNow(date, now) {
			break
		}
	}
	return date.Format("20060102"), nil
}

func processCommaSeparatedParts(str string, minValue, maxValue int) ([]int, error) {
	if str == "" {
		return []int{}, errors.New("Bad repeat parameter value")
	}
	parts := strings.Split(str, ",")
	slice := make([]int, 0, len(parts))
	for _, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return []int{}, errors.New("Bad repeat parameter value")
		}
		if value < minValue || value > maxValue {
			return []int{}, errors.New("Bad repeat parameter value")
		}
		slice = append(slice, value)
	}
	return slice, nil
}

func processWeeklyRepeat(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	wstr := strings.Split(repeat, " ")
	if len(wstr) != 2 {
		return "", errors.New("Bad repeat parameter value")
	}
	weekdaysStr := strings.Split(wstr[1], ",")
	if len(weekdaysStr) == 0 {
		return "", errors.New("Bad repeat parameter value")
	}
	weekdays, err := processCommaSeparatedParts(wstr[1], 1, 7)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(weekdays); i++ {
		if weekdays[i] == 7 {
			weekdays[i] = 0
		}
	}
	for {
		date = date.AddDate(0, 0, 1)
		currentWeekday := int(date.Weekday())
		if AfterNow(date, now) && slices.Contains(weekdays, currentWeekday) {
			break
		}
	}
	return date.Format("20060102"), nil
}

func daysInMonth(date time.Time) int {
	t := time.Date(date.Year(), date.Month(), 32, 0, 0, 0, 0, time.UTC)
	return 32 - t.Day()
}

func processMonthlyRepeat(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	mstr := strings.Split(repeat, " ")
	if len(mstr) < 2 && len(mstr) > 3 {
		return "", errors.New("Bad repeat parameter value")
	}

	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	days, err := processCommaSeparatedParts(mstr[1], -2, 31)
	if err != nil {
		return "", errors.New("Bad repeat parameter value")
	}
	if len(mstr) == 3 {
		months, err = processCommaSeparatedParts(mstr[2], 1, 12)
		if err != nil {
			return "", errors.New("Bad repeat parameter value")
		}
	}
	for {
		date = date.AddDate(0, 0, 1)
		dayInt := int(date.Day())
		monthInt := int(date.Month())
		if AfterNow(date, now) && slices.Contains(months, monthInt) && (slices.Contains(days, dayInt) || slices.Contains(days, dayInt-daysInMonth(date)-1)) {
			break
		}
	}
	return date.Format("20060102"), nil
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", errors.New("Bad repeat parameter value")
	}
	switch string(repeat[0]) {
	case "d":
		return processDailyRepeat(now, dstart, repeat)
	case "y":
		return processYearlyRepeat(now, dstart, repeat)
	case "w":
		return processWeeklyRepeat(now, dstart, repeat)
	case "m":
		return processMonthlyRepeat(now, dstart, repeat)
	}
	return "", errors.New("Bad repeat parameter value")
}
