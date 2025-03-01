// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteTask(ctx context.Context, id int64) error
	GetTask(ctx context.Context, id int64) (Task, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListTasks(ctx context.Context, arg ListTasksParams) ([]Task, error)
	UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error)
}

var _ Querier = (*Queries)(nil)
