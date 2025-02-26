-- name: CreateTask :one
INSERT INTO task (
        title,
        description,
        status,
        priority,
        executor
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;