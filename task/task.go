package task

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
	DueDate   time.Time
}

type ToDoList struct {
	db *sql.DB
}

func NewToDoList(db *sql.DB) *ToDoList {
	return &ToDoList{
		db: db,
	}
}

func (list *ToDoList) AddTask(title string, dueDate time.Time) error {
	_, err := list.db.Exec("INSERT INTO tasks (title, completed, due_date) VALUES (?, ?, ?)",
		title, false, dueDate)
	return err
}

func (list *ToDoList) CompleteTask(taskID int) error {
	query := "UPDATE tasks SET completed = ? WHERE id = ?"

	_, err := list.db.Exec(query, true, taskID)
	if err != nil {
		return err
	}

	return nil
}

func (list *ToDoList) RemoveTask(taskID int) error {
	query := "DELETE FROM tasks WHERE id = ?"

	_, err := list.db.Exec(query, taskID)
	if err != nil {
		return err
	}

	return nil
}

func (list *ToDoList) ListTasks() ([]Task, error) {
	rows, err := list.db.Query("SELECT id, title, completed, due_date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var dueDateStr string
		err := rows.Scan(&task.ID, &task.Title, &task.Completed, &dueDateStr)
		if err != nil {
			return nil, err
		}
		task.DueDate, _ = time.Parse("2006-01-02 15:04:05", dueDateStr)
		tasks = append(tasks, task)
	}

	return tasks, nil
}
