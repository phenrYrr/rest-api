package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yggdrasiI1/rest-api/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Author sql.NullString
	Body   sql.NullString
}

func ConvertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow

	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug , body, author
		FROM comments
		WHERE id = $1`,
		uuid,
	)

	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid")
	}

	return comment.Comment{}, nil
}
