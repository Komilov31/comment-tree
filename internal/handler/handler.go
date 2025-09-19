package handler

import (
	"github.com/Komilov31/comment-tree/internal/dto"
	"github.com/Komilov31/comment-tree/internal/model"
)

type CommentService interface {
	GetAllComments() ([]*model.Comment, error)
	GetCommentsById(int) ([]*model.Comment, error)
	GetCommentsPaginated(dto.CommentsPagination) ([]*model.Comment, error)
	GetCommentsByTextSearch(string) ([]*model.Comment, error)
	CreateComment(dto.CreateComment) (*dto.CreateComment, error)
	DeleteCommentById(int) error
}

type Handler struct {
	service CommentService
}

func New(service CommentService) *Handler {
	return &Handler{
		service: service,
	}
}
