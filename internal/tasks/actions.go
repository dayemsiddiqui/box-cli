package tasks

import "github.com/dayemsiddiqui/box-cli/internal/db"


func CreateTaskAction(name string, description string) error {
	db := db.EnsureDB()
	defer db.Close()
	taskRepository := NewTaskRepository(db)
	task := NewTask(name, &description)
	return taskRepository.SaveTask(task)
}


func GetTaskAction(id string) (*Task, error) {
	db := db.EnsureDB()
	defer db.Close()
	taskRepository := NewTaskRepository(db)
	return taskRepository.GetTask(id)
}

func GetAllTasksAction() ([]*Task, error) {
	db := db.EnsureDB()
	defer db.Close()
	taskRepository := NewTaskRepository(db)
	return taskRepository.GetAllTasks()
}