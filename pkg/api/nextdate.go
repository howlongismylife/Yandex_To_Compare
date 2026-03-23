package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {

	case "d":
		if len(parts) != 2 {
			return "", errors.New("invalid repeat format")
		}

		interval, err := strconv.Atoi(parts[1])
		if err != nil || interval < 1 || interval > 400 {
			return "", errors.New("invalid day interval")
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}

		return date.Format(dateFormat), nil

	case "y":
		if len(parts) != 1 {
			return "", errors.New("invalid repeat format")
		}

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}

		return date.Format(dateFormat), nil

	default:
		return "", errors.New("unsupported repeat format")
	}
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeat := r.FormValue("repeat")

	if dateStr == "" {
		writeJSON(w, map[string]any{"error": "date is required"})
		return
	}

	now := time.Now()
	var err error

	if nowStr != "" {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			writeJSON(w, map[string]any{"error": "invalid now"})
			return
		}
	}

	next, err := NextDate(now, dateStr, repeat)
	if err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, map[string]string{
		"date": next,
	})
}
