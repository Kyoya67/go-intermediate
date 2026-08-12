package controllers_test

import (
	"testing"

	"github.com/Kyoya67/go-intermediate/controllers"
	"github.com/Kyoya67/go-intermediate/controllers/testdata"

	_ "github.com/go-sql-driver/mysql"
)

var articleController *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	articleController = controllers.NewArticleController(ser)

	m.Run()
}
