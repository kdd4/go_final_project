package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/kdd4/go_final_project/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}

	t, err := time.Parse(dateFormat, task.Date)

	if err != nil {
		return err
	}

	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)

		if err != nil {
			return err
		}
	}

	if afterNow(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format(dateFormat)
		} else {
			task.Date = next
		}
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, req *http.Request) {
	var (
		task db.Task
		buf bytes.Buffer 
	)

	_, err := buf.ReadFrom(req.Body)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}

	if task.Title == "" {
		writeJsonError(w, "title is empty")
		return
	}

	err = checkDate(&task)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}

	id, err := db.AddTask(&task)

	result := make(map[string]string) 

	if err != nil {
		result["error"] = err.Error()
	} else {
		result["id"] = strconv.FormatInt(id, 10)
	}

	writeJson(w, result, http.StatusOK)
}