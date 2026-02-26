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

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	if string(repeat[0]) == "d" {
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

	} else if repeat == "y" {
		for {
			date = date.AddDate(1, 0, 0)
			if AfterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil
	} else if string(repeat[0]) == "w" {
		wstr := strings.Split(repeat, " ")
		if len(wstr) != 2 {
			return "", errors.New("Bad repeat parameter value")
		}
		daystr := strings.Split(wstr[1], ",")
		if len(daystr) == 0 {
			return "", errors.New("Bad repeat parameter value")
		}
		days := make([]int, 0, len(daystr))
		for _, dstr := range daystr {
			day, err := strconv.Atoi(dstr)
			if err != nil {
				return "", errors.New("Bad repeat parameter value")
			}
			if day < 1 || day > 7 {
				return "", errors.New("Bad repeat parameter value")
			}
			if day == 7 {
				day = 0
			}
			days = append(days, day)
		}
		for {
			date = date.AddDate(0, 0, 1)
			dayInt := int(date.Weekday())
			if AfterNow(date, now) && slices.Contains(days, dayInt) {
				break
			}
		}
		return date.Format("20060102"), nil
	} else if string(repeat[0]) == "m" {
		mstr := strings.Split(repeat, " ")
		if len(mstr) < 2 && len(mstr) > 3 {
			return "", errors.New("Bad repeat parameter value")
		}
		if len(mstr) == 2 {
			weekdaysstr := strings.Split(mstr[1], ",")
			if len(weekdaysstr) == 0 {
				return "", errors.New("Bad repeat parameter value")
			}
			weekdays := make([]int, 0, len(weekdaysstr))
			for _, weekdaystr := range weekdaysstr {
				weekday, err := strconv.Atoi(weekdaystr)
				if err != nil {
					return "", errors.New("Bad repeat parameter value")
				}
				if weekday < -2 || weekday > 31 {
					return "", errors.New("Bad repeat parameter value")
				}
				weekdays = append(weekdays, weekday)
			}
			for {
				date = date.AddDate(0, 0, 1)
				dayInt := int(date.Day())
				t := time.Date(date.Year(), date.Month(), 32, 0, 0, 0, 0, time.UTC)
				daysInMonth := 32 - t.Day()
				if AfterNow(date, now) && (slices.Contains(weekdays, dayInt) || slices.Contains(weekdays, dayInt-daysInMonth-1)) {
					break
				}
			}
			return date.Format("20060102"), nil
		}

		if len(mstr) == 3 {
			weekdaysstr := strings.Split(mstr[1], ",")
			if len(weekdaysstr) == 0 {
				return "", errors.New("Bad repeat parameter value")
			}
			weekdays := make([]int, 0, len(weekdaysstr))
			for _, weekdaystr := range weekdaysstr {
				weekday, err := strconv.Atoi(weekdaystr)
				if err != nil {
					return "", errors.New("Bad repeat parameter value")
				}
				if weekday < -2 || weekday > 31 {
					return "", errors.New("Bad repeat parameter value")
				}
				weekdays = append(weekdays, weekday)
			}
			monthsstr := strings.Split(mstr[2], ",")
			if len(monthsstr) == 0 {
				return "", errors.New("Bad repeat parameter value")
			}
			months := make([]int, 0, len(monthsstr))
			for _, monthstr := range monthsstr {
				month, err := strconv.Atoi(monthstr)
				if err != nil {
					return "", errors.New("Bad repeat parameter value")
				}
				if month < 1 || month > 12 {
					return "", errors.New("Bad repeat parameter value")
				}
				months = append(months, month)
			}
			for {
				date = date.AddDate(0, 0, 1)
				dayInt := int(date.Day())
				monthInt := int(date.Month())
				t := time.Date(date.Year(), date.Month(), 32, 0, 0, 0, 0, time.UTC)
				daysInMonth := 32 - t.Day()
				if AfterNow(date, now) && slices.Contains(months, monthInt) && (slices.Contains(weekdays, dayInt) || slices.Contains(weekdays, dayInt-daysInMonth-1)) {
					break
				}
			}
			return date.Format("20060102"), nil
		}

	}
	return "", errors.New("Bad repeat parameter value")
}
