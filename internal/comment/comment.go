package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {

	cmt, err := s.Store.GetComment(ctx, id)

	ctx = context.WithValue(ctx, "request_id", "unique-string")

	fmt.Println(ctx.Value("request_id"))

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, updatedCmt Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Errorf("error updating comment")
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return ErrNotImplemented
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		return Comment{}, err
	}

	return insertedCmt, nil
}
