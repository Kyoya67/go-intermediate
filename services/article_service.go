package services

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/Kyoya67/go-intermediate/apperrors"
	"github.com/Kyoya67/go-intermediate/models"
	"github.com/Kyoya67/go-intermediate/repositories"
)

func (s *MyAppService) GetArticleDetailService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	var wg sync.WaitGroup
	wg.Add(2)

	var amu sync.Mutex
	var cmu sync.Mutex

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		newArticle, err := repositories.SelectArticleDetail(s.db, articleID)
		amu.Lock()
		article, articleGetErr = newArticle, err
		amu.Unlock()
	}(s.db, articleID)

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		newCommentList, err := repositories.SelectCommentList(s.db, articleID)
		cmu.Lock()
		commentList, commentGetErr = newCommentList, err
		cmu.Unlock()
	}(s.db, articleID)

	wg.Wait()

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = commentList

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return articleList, nil
}

func (s *MyAppService) PostNiceService(articleID int) (models.Article, error) {
	updatedArticle, err := repositories.UpdateNiceNum(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	return updatedArticle, nil
}
