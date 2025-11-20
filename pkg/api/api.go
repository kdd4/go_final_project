package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const dateFormat = "20060102"

func Init(r chi.Router) {
	r.Get("/nextdate", nextDateHandler)
	r.Get("/tasks", tasksHandler)
	r.Get("/task", getTaskHandler)

	r.Post("/task", addTaskHandler)
	r.Put("/task", putTaskHandler)

	r.Post("/task/done", doneTaskHandler)
	r.Delete("/task", deleteTaskHandler)
}


func writeJson(w http.ResponseWriter, data any, code int) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		writeJsonError(w, err.Error())
		return
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(jsonData)
}

func writeJsonError(w http.ResponseWriter, err string) {
	writeJson(
		w, 
		map[string]string{
			"error": err,
		}, 
		http.StatusBadRequest,
	)
}