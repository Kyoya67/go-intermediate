package main

import (
	"database/sql"
	"dbsample/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	articleID := 10
	const sqlStr = `
	SELECT title, contents, username, nice, created_at
	FROM articles
	WHERE article_id = ?
	`
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdAt sql.NullTime
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)

		if createdAt.Valid {
			article.CreatedAt = createdAt.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
