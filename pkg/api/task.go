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
		writeJsonError(w, "id is not specified", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeJsonError(w, "task not found", http.StatusBadRequest)
			return
		}
		writeJsonError(w, err.Error(), http.StatusInternalServerError)
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
		writeJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)

	if err != nil {
		writeJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJsonError(w, "title is empty", http.StatusBadRequest)
		return
	}

	err = checkDate(&task)

	if err != nil {
		writeJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)

	if err != nil {
		writeJson(
			w, 
			map[string]any{
				"error": err.Error(),
			}, 
			http.StatusOK,
		)
		return
	}
	
	writeJson(
		w, 
		map[string]any{
			"id": task.ID,
		}, 
		http.StatusOK,
	)
}

func deleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		writeJsonError(w, "id is not specified", http.StatusBadRequest)
		return
	}

	err := db.DeleteTask(id)

	if err != nil {
		writeJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJson(w, struct{}{}, http.StatusOK)
}

func doneTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		writeJsonError(w, "id is not specified", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		writeJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(task.ID)

		if err != nil {
			writeJsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJson(w, struct{}{}, http.StatusOK)
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)

	if err != nil {
		writeJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateDate(next, task.ID)

	if err != nil {
		writeJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(w, struct{}{}, http.StatusOK)
}
