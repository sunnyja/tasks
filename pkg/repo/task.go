package task

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Task struct {
	ID          int
	Title	    string
	Description string
	Status		int
	CreatedAt  int64
	ClosedAt   int64
	UpdatedAt	int64
	AuthorId    int
	AssignedId  int
}

type TaskStorage struct {
	pool *pgxpool.Pool
}


func NewTaskStorage(pool *pgxpool.Pool) *TaskStorage {
	return &TaskStorage{
		pool: pool,
	}
}

//вернет все задачи / задачи по id / задачи по автору
func (ts *TaskStorage) GetTasks(ctx context.Context, taskID int, authorID int) ([]Task, error) {
	var tasks []Task
	const query = `
		SELECT 
			id,
			title,
			description,
			created_at,
			author_id,
			assigned_id,
			status
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) 
		AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;`

	rows, err := ts.pool.Query(ctx, query, taskID, authorID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.CreatedAt,
			&t.AuthorId,
			&t.AssignedId,
			&t.Status,
		)
		if err != nil {
			return nil, err
		}
		
		tasks = append(tasks, t)

	}
	return tasks, rows.Err()
}


// NewTask создаёт новую задачу и возвращает её id
func (ts *TaskStorage) NewTask(ctx context.Context, t Task) (int, error) {
	var id int
	const query = `INSERT INTO tasks (title, description, author_id, assigned_id) VALUES ($1, $2, $3, $4) RETURNING id;`

	err := ts.pool.QueryRow(ctx, query, t.Title, t.Description, t.AuthorId, t.AssignedId).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}

	return id, nil
}

//обновление задачи
func (ts *TaskStorage) UpdateTask(ctx context.Context, taskID int, t Task) error {
	const query = `
		UPDATE tasks
		SET 
			title = $2, 
			description = $3,
			author_id = $4,
			assigned_id = $5,
			status = $6,
			updated_at = NOW()
		WHERE id = $1;`

	_, err := ts.pool.Exec(ctx, query, taskID, t.Title, t.Description, t.AuthorId, t.AssignedId, t.Status)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	
	return nil
}

//удаление задачи
func (ts *TaskStorage) DeleteTask(ctx context.Context, taskID int) error {
	const query = `DELETE FROM tasks WHERE id = $1`
	
	_, err := ts.pool.Exec(ctx, query, taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	
	return nil
}