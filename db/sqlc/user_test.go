package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/hisshihi/todo-list-go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomString(6),
		HashedPassword: util.RandomString(12),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.Username)
	require.NotZero(t, user.HashedPassword)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
