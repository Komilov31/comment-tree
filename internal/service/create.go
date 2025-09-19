package service

import (
	"github.com/Komilov31/comment-tree/internal/dto"
)

func (s *Service) CreateComment(comment dto.CreateComment) (*dto.CreateComment, error) {
	return s.storage.CreateComment(comment)
}
