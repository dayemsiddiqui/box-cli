package tasks

import (
	"database/sql"
)


type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) SaveTask(task *Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (id, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", task.ID, task.Title, task.Description, task.CreatedAt, task.UpdatedAt)
	return err
}

func (r *TaskRepository) GetTask(id string) (*Task, error) {
	row := r.db.QueryRow("SELECT id, name, description, created_at, updated_at FROM tasks WHERE id = ?", id)
	task := &Task{}
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) GetAllTasks() ([]*Task, error) {
	rows, err := r.db.Query("SELECT id, name, description, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []*Task{}
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}