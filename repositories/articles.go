package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Kyoya67/go-intermediate/models"
)

const (
	articleNumPerPage = 5
)

// 新規投稿をDBにinsertする関数
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	INSERT INTO articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

// 投稿一覧をDBから取得する関数
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		SELECT article_id, title, contents, username, nice
		FROM articles
		LIMIT ? OFFSET ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
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

// 投稿IDを指定して、記事データを取得する関数
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

func UpdateNiceNum(db *sql.DB, articleID int) (models.Article, error) {
	const sqlUpdateNice = `UPDATE articles SET nice = nice + 1 WHERE article_id = ?`

	result, err := db.Exec(sqlUpdateNice, articleID)
	if err != nil {
		return models.Article{}, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return models.Article{}, err
	}

	if affected == 0 {
		return models.Article{}, fmt.Errorf("no rows affected")
	}

	selectedArticle, err := SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	return selectedArticle, nil
}
