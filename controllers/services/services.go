package services

import "github.com/Kyoya67/go-intermediate/models"

type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleDetailService(id int) (models.Article, error)
	PostNiceService(articleID int) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
	GetCommentListService(articleID int) ([]models.Comment, error)
}
