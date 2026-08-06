package services

import (
	"github.com/Kyoya67/go-intermediate/models"
	"github.com/Kyoya67/go-intermediate/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}

func GetCommentListService(articleID int) ([]models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return nil, err
	}

	return commentList, nil
}
