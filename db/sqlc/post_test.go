package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vule96/my-blog-be/util"
)

func createRandomPost(t *testing.T) Posts {
	arg := CreatePostParams{
		ID:      util.GenUUID(),
		UserID:  util.GenUUID(),
		Content: util.RandomString(50),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.Content, post.Content)

	require.NotEmpty(t, post.ID)
	require.NotZero(t, post.IsActive)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := createRandomPost(t)
	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.Content, post2.Content)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, post1.IsActive, post2.IsActive)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
	require.WithinDuration(t, post1.UpdatedAt, post2.UpdatedAt, time.Second)
}

func TestUpdatePost(t *testing.T) {
	post1 := createRandomPost(t)

	arg := UpdatePostParams{
		ID:      post1.ID,
		Content: util.RandomString(25),
	}

	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, arg.Content, post2.Content)
	require.Equal(t, post1.IsActive, post2.IsActive)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
	require.WithinDuration(t, post1.UpdatedAt, post2.UpdatedAt, time.Second)
}

func TestDeletePost(t *testing.T) {
	post1 := createRandomPost(t)
	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}

func TestListPosts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}

	arg := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}

	posts, err := testQueries.ListPosts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, posts, 5)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}
