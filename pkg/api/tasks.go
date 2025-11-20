package api

import (
	"net/http"

	"github.com/kdd4/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, req *http.Request) {
	tasks, err := db.Tasks(50)

	if err != nil {
		writeJsonError(w, err.Error())
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