package funcforhand

import (
	"fmt"

	m "github.com/vova1001/Test_task_manager/model"
)

var (
	sliceTasks []m.Task
	NewID      = 1
)

func GetTasks(logChan chan<- string) []m.Task {
	logChan <- "all tasks requested"
	return sliceTasks
}

func GetTaskId(id int, logChan chan<- string) (m.Task, error) {
	for _, task := range sliceTasks {
		if task.ID == id {
			logChan <- "found a task with the required ID"
			return task, nil
		}
	}
	logChan <- "I didn't find a task with such an ID"
	return m.Task{}, fmt.Errorf("Invalid ID")
}

func PostTask(task m.Task, logChan chan<- string) m.Task {
	task.ID = NewID
	NewID++
	sliceTasks = append(sliceTasks, task)
	logChan <- "task added successfully"
	return task
}
