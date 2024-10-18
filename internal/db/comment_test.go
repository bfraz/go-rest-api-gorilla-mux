//go:build integration
// +build integration

package db

import (
	"context"
	"go-rest-api-course-v2/internal/comment"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		t.Run("test create commment", func(t *testing.T) {
			db, err := NewDatabase()
			assert.NoError(t, err)

			cmt, err := db.PostComment(context.Background(), comment.Comment{
				Slug:   "slug",
				Body:   "body",
				Author: "author",
			})
			assert.NoError(t, err)

			newCmt, err := db.GetComment(context.Background(), cmt.ID)
			assert.NoError(t, err)
			assert.Equal(t, "slug", newCmt.Slug)
			assert.Equal(t, "body", newCmt.Body)
			assert.Equal(t, "author", newCmt.Author)
		})

		t.Run("test delete commment", func(t *testing.T) {
			db, err := NewDatabase()
			assert.NoError(t, err)

			cmt, err := db.PostComment(context.Background(), comment.Comment{
				Slug:   "new-slug",
				Body:   "new-body",
				Author: "new-author",
			})
			assert.NoError(t, err)

			err = db.DeleteComment(context.Background(), cmt.ID)
			assert.NoError(t, err)

			_, err = db.GetComment(context.Background(), cmt.ID)
			assert.Error(t, err)
		})
	})
}
