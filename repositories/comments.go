package repositories

import (
	"database/sql"
	"fmt"

	"github.com/kyoya67/go-intermediate/models"
)

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
