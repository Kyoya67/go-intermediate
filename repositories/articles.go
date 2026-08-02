package repositories

import (
	"database/sql"
	"fmt"

	"github.com/kyoya67/myapi/models"
)

const articleNumPerpage = 5

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqltr = `
	INSERT INTO articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqltr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		SELECT article_id, title, contents, username, nice
		FROM articles
		LIMIT ? OFFSET ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerpage, (page-1)*articleNumPerpage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		SELECT *
		FROM articles
		WHERE article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		INSERT INTO comments (article_id, message, created_at) values
		(?, ?, now());
	`

	var newComment models.Comment
	newComment.CommentID, newComment.ArticleID, newComment.Message, newComment.CreatedAt = comment.CommentID, comment.ArticleID, comment.Message, comment.CreatedAt

	_, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		fmt.Println(err)
	}

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	commentList := make([]models.Comment, 0)
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if err != nil {
			return nil, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentList = append(commentList, comment)
	}
	return commentList, nil
}
