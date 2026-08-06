package services

import (
	"github.com/Kyoya67/go-intermediate/models"
	"github.com/Kyoya67/go-intermediate/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connect()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = commentList

	return article, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connect()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func PostNiceService(articleID int) (models.Article, error) {
	db, err := connect()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.UpdateNiceNum(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}
