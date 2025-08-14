package handler

import (
	"encoding/json"
	"net/http"

	f "github.com/vova1001/Test_task_manager/funcforhand"
	m "github.com/vova1001/Test_task_manager/model"
)

func GetAllHendler(w http.ResponseWriter, r *http.Request, logChan chan<- string) {
	Result := f.GetTasks(logChan)
	logChan <- "Return all tasks"
	json.NewEncoder(w).Encode(Result)
}

func GetIdHendler(w http.ResponseWriter, r *http.Request, id int, logChan chan<- string) {
	Result, err := f.GetTaskId(id, logChan)
	if err != nil {
		http.Error(w, "Task not found", 404)
		logChan <- "Task not found"
		return
	}
	logChan <- "Return taskId"
	json.NewEncoder(w).Encode(Result)

}

func PostHendler(w http.ResponseWriter, r *http.Request, logChan chan<- string) {
	var task m.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid body JSON", 404)
		logChan <- "Invalid body JSON"
		return
	}

	result := f.PostTask(task, logChan)
	logChan <- "successful post"
	json.NewEncoder(w).Encode(result)

}
