package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/hisshihi/todo-list-go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTask(t *testing.T) Task {
	user1 := CreateUserParams{
		Username: util.RandomString(6),
		HashedPassword: "123456",
	}

	user, err := testQueries.CreateUser(context.Background(), user1)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	arg := CreateTaskParams{
		Title: "Сходить в магазин",
		Description: "Купить молоко, хлеб, колбасу, сыр",
		Status: "Активно",
		Priority: "Срочно",
		Executor: user.Username,
	}

	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.Status, task.Status)
	require.Equal(t, arg.Priority, task.Priority)
	require.Equal(t, arg.Executor, task.Executor)

	require.NotZero(t, task.ID)
	require.NotZero(t, task.CreatedAt)

	return task
}

func TestCreateTask(t *testing.T) {
	CreateRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := CreateRandomTask(t)
	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, task1.Title, task2.Title)
	require.Equal(t, task1.Description, task2.Description)
	require.Equal(t, task1.Status, task2.Status)
	require.Equal(t, task1.Priority, task2.Priority)
	require.Equal(t, task1.Executor, task2.Executor)

	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second)
}

func TestListTasks(t *testing.T) {
	for range 5 {
		CreateRandomTask(t)
	}

	arg := ListTasksParams{
		Limit:  5,
		Offset: 0,
	}

	tasks, err := testQueries.ListTasks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)
	require.Len(t, tasks, int(arg.Limit))

	for _, task := range tasks {
		require.NotZero(t, task.ID)
		require.NotZero(t, task.CreatedAt)
	}

	t.Run("EmptyTasksList", func(t *testing.T) {
		// Тестируем случай, когда список задач пуст
		argEmpty := ListTasksParams{
			Limit:  5,
			Offset: 999999999999999999,
		}
		tasksEmpty, err := testQueries.ListTasks(context.Background(), argEmpty)
		require.NoError(t, err)
		require.Empty(t, tasksEmpty)
		require.Len(t, tasksEmpty, 0)
	})
}

func TestUpdateTask(t *testing.T) {
	task1 := CreateRandomTask(t)

	arg := UpdateTaskParams{
		ID: task1.ID,
		Title: "Сходить в продуктовый магазин",
		Description: "Купить молоко, печенье, сыр",
		Status: "Выполнено",
		Priority: "Средне",
		Executor: task1.Executor,
	}

	task2, err := testQueries.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, arg.Title, task2.Title)
	require.Equal(t, arg.Description, task2.Description)
	require.Equal(t, arg.Status, task2.Status)
	require.Equal(t, arg.Priority, task2.Priority)
	require.Equal(t, arg.Executor, task2.Executor)

	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second)
}

func TestDeleteTask(t *testing.T) {
	task1 := CreateRandomTask(t)
	err := testQueries.DeleteTask(context.Background(), task1.ID)
	require.NoError(t, err)

	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task2)
}