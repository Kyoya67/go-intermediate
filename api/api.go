package api

import (
	"database/sql"
	"net/http"

	"github.com/Kyoya67/go-intermediate/api/middlewares"
	"github.com/Kyoya67/go-intermediate/controllers"
	"github.com/Kyoya67/go-intermediate/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	service := services.NewMyAppService(db)
	articleController := controllers.NewArticleController(service)
	commentController := controllers.NewCommentController(service)

	r := mux.NewRouter()

	r.HandleFunc("/article", articleController.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", articleController.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", articleController.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", articleController.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", commentController.PostCommentHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment/list", commentController.GetCommentListHandler).Methods(http.MethodGet)

	r.Use(middlewares.LoggingMiddleware)

	return r
}
