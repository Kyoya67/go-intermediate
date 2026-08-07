package services

import (
	"github.com/Kyoya67/go-intermediate/models"
	"github.com/Kyoya67/go-intermediate/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}

func (s *MyAppService) GetCommentListService(articleID int) ([]models.Comment, error) {

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return nil, err
	}

	return commentList, nil
}
