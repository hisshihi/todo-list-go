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

-- name: GetTask :one
SELECT * FROM task
WHERE id = $1;

-- name: ListTasks :many
SELECT * FROM task
ORDER BY created_at DESC;

-- name: UpdateTask :one
UPDATE task
SET title = $2, description = $3, status = $4, priority = $5, executor = $6
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM task
WHERE id = $1;