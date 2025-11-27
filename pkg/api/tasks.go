package api

import (
	"net/http"

	"github.com/kdd4/go_final_project/pkg/db"
)

const tasksLimit = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, req *http.Request) {
	tasks, err := db.Tasks(tasksLimit)

	if err != nil {
		writeJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(
		w, 
		TasksResp{
			Tasks: tasks,
		}, 
		http.StatusOK,
	)
}