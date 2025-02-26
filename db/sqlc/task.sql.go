// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: task.sql

package sqlc

import (
	"context"
)

const createTask = `-- name: CreateTask :one
INSERT INTO task (
        title,
        description,
        status,
        priority,
        executor
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, description, status, created_at, updated_at, priority, executor
`

type CreateTaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Priority    string `json:"priority"`
	Executor    string `json:"executor"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.Priority,
		arg.Executor,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Priority,
		&i.Executor,
	)
	return i, err
}
