package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func afterNow(date, now time.Time) bool {
	y1, m1, d1 := date.Date()
	y2, m2, d2 := now.Date()

	if y1 != y2 {
		return y1 > y2
	}

	if m1 != m2 {
		return m1 > m2
	}

	return d1 > d2
} 

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if dstart == "" || repeat == "" {
		return "", errors.New("dstart or/and repear empty")
	}

	date, err := time.Parse(dateFormat, dstart);

	if err != nil {
		return "", fmt.Errorf("error while parsing string %q: %w", dstart, err)
	}

	splitedRepeat := strings.Split(repeat, " ")

	var (
		dayInterval = 0
		yearInterval = 0
	)

	switch len(splitedRepeat) {
	case 1:
		if splitedRepeat[0] != "y" {
			return "", errors.New("wrong repeat string")
		}
		yearInterval = 1

	case 2:
		if splitedRepeat[0] != "d" {
			return "", errors.New("wrong repeat string")
		}

		days, err := strconv.ParseInt(splitedRepeat[1], 10, 32)
		
		if err != nil {
			return "", errors.New("wrong format of days count in repeat string")
		}

		dayInterval = int(days)

		if dayInterval > 400 {
			return "", errors.New("days count more than 400")
		}

	default:
		return "", errors.New("wrong len of repeat string") 
	}


	for {
		date = date.AddDate(yearInterval, 0, dayInterval)

		if afterNow(date, now) {
			break
		}
	}

	nextDate := date.Format(dateFormat)

	return nextDate, nil
}

func nextDateHandler(w http.ResponseWriter, req *http.Request) {
	now := req.FormValue("now")
	date := req.FormValue("date")
	repeat := req.FormValue("repeat")

	if date == "" || repeat == "" {
		http.Error(w, "dstart or/and repear empty", http.StatusBadRequest)
		return
	}

	var (
		nowDate time.Time
		err error
	);

	if now != "" {
		nowDate, err = time.Parse(dateFormat, now)

		if err != nil {
			http.Error(w, "now date parsing error", http.StatusBadRequest)
			return
		}
	} else {
		nowDate = time.Now()
	}

	next, err := NextDate(nowDate, date, repeat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, next)
}