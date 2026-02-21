package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func afterNow(date, now time.Time) bool {
	return date.After(now)
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
			if afterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil

	} else if repeat == "y" {
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil
	}
	return "", errors.New("Bad repeat parameter value")
}

func nextDateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		nowstr := req.FormValue("now")
		now, _ := time.Parse("20060102", nowstr)
		date := req.FormValue("date")
		repeat := req.FormValue("repeat")
		if repeat == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result, err := NextDate(now, date, repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}
