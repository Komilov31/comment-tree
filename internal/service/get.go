package service

import (
	"github.com/Komilov31/comment-tree/internal/dto"
	"github.com/Komilov31/comment-tree/internal/model"
)

func (s *Service) GetAllComments() ([]*model.Comment, error) {
	return s.storage.GetAllComments()
}

func (s *Service) GetCommentsById(id int) ([]*model.Comment, error) {
	return s.storage.GetCommentsById(id)
}

func (s *Service) GetCommentsPaginated(config dto.CommentsPagination) ([]*model.Comment, error) {
	return s.storage.GetCommentsPaginated(config)
}

func (s *Service) GetCommentsByTextSearch(text string) ([]*model.Comment, error) {
	return s.storage.GetCommentsByTextSearch(text)
}
