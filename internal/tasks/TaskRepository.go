package tasks

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)


type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) SaveTask(task *Task) error {
	query, args, err := sq.Insert("tasks").Columns("id", "name", "description", "created_at", "updated_at").Values(task.ID, task.Title, task.Description, task.CreatedAt, task.UpdatedAt).ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func (r *TaskRepository) GetTask(id string) (*Task, error) {
	query, args, err := sq.Select("id", "name", "description", "created_at", "updated_at").From("tasks").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRow(query, args...)
	task := &Task{}
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) GetAllTasks() ([]*Task, error) {
	query, args, err := sq.Select("id", "name", "description", "created_at", "updated_at").From("tasks").ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(query, args...)
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
		// If the description is null, set it to an empty string
		if task.Description == "" {
			task.Description = ""
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}


func (r *TaskRepository) DeleteAllTasks() error {
	query, args, err := sq.Delete("tasks").ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}