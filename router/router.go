package router

import (
	"net/http"
	"strconv"
	"strings"

	h "github.com/vova1001/Test_task_manager/handler"
)

func RegisterRouterTask(logChan chan<- string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathTask := strings.Trim(r.URL.Path, "/")
		switch pathTask {
		case "":
			if r.Method == http.MethodGet {
				logChan <- "GetAll"
				h.GetAllHendler(w, r, logChan)
				return
			}
			if r.Method == http.MethodPost {
				logChan <- "Post"
				h.PostHendler(w, r, logChan)
				return
			}
		default:
			id, err := strconv.Atoi(pathTask)
			if err != nil {
				http.Error(w, "Invalid ID", 404)
				logChan <- "Invalid ID"
				return
			}
			if r.Method == http.MethodGet {
				logChan <- "GetId"
				h.GetIdHendler(w, r, id, logChan)
				return
			} else {
				http.Error(w, "Invalid method", 405)
				logChan <- "Invalid method taskId"
			}
		}
	}
}
