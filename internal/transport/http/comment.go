package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yggdrasiI1/rest-api/internal/comment"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type Response struct {
	Message string
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Body   string `json:"body" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func ConvertPostCommentRequestToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Body:   c.Body,
		Author: c.Author,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(cmt)

	if err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}

	convertedComment := ConvertPostCommentRequestToComment(cmt)

	postedComment, err := h.Service.PostComment(r.Context(), convertedComment)

	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		panic(err)
	}

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	commentId := vars["id"]

	if commentId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), commentId)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}
