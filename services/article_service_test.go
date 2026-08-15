package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/Kyoya67/go-intermediate/services"

	_ "github.com/go-sql-driver/mysql"
)

var articleService *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	articleService = services.NewMyAppService(db)

	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := articleService.GetArticleDetailService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
