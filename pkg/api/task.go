package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kdd4/go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		writeJsonError(w, "id is not specified")
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeJsonError(w, "task not found")
			return
		}
		writeJsonError(w, err.Error())
	}

	writeJson(w, task, http.StatusOK)
}

func putTaskHandler(w http.ResponseWriter, req *http.Request) {
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

	err = db.UpdateTask(&task)

	result := make(map[string]string) 

	if err != nil {
		result["error"] = err.Error()
	} else {
		result["id"] = task.ID
	}

	writeJson(w, result, http.StatusOK)
}

func deleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		writeJsonError(w, "id is not specified")
		return
	}

	err := db.DeleteTask(id)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}

	writeJson(w, struct{}{}, http.StatusOK)
}

func doneTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		writeJsonError(w, "id is not specified")
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(task.ID)

		if err != nil {
			writeJsonError(w, err.Error())
			return
		}

		writeJson(w, struct{}{}, http.StatusOK)
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}

	err = db.UpdateDate(next, task.ID)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}

	writeJson(w, struct{}{}, http.StatusOK)
}
