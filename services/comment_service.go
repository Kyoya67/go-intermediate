package services

import (
	"github.com/Kyoya67/go-intermediate/apperrors"
	"github.com/Kyoya67/go-intermediate/models"
	"github.com/Kyoya67/go-intermediate/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}

func (s *MyAppService) GetCommentListService(articleID int) ([]models.Comment, error) {
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	return commentList, nil
}
