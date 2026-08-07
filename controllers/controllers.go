package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Kyoya67/go-intermediate/models"
	"github.com/gorilla/mux"

	"github.com/Kyoya67/go-intermediate/services"
)

type MyAppController struct {
	service *services.MyAppService
}

func NewMyAppController(service *services.MyAppService) *MyAppController {
	return &MyAppController{service: service}
}

// POST /article のハンドラ
func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail to post article\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// GET /article/list のハンドラ
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail to get article list\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id} のハンドラ
func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 暫定でこれを追加することで
	// 「変数articleIDが使われていない」というコンパイルエラーを回避
	log.Println(articleID)

	article, err := c.service.GetArticleDetailService(articleID)
	if err != nil {
		http.Error(w, "fail to get article detail\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /article/nice のハンドラ
func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle.ID)
	if err != nil {
		http.Error(w, "fail to post nice\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// POST /comment のハンドラ
func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to post comment\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}

func (c *MyAppController) GetCommentListHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	commentList, err := c.service.GetCommentListService(articleID)
	if err != nil {
		http.Error(w, "fail to get comment list\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(commentList)
}
