package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kyoya67/go-intermediate/models"
	"github.com/gorilla/mux"

	"github.com/Kyoya67/go-intermediate/apperrors"
	"github.com/Kyoya67/go-intermediate/controllers/services"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to post comment\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}

// GET /comment/list のハンドラ
func (c *CommentController) GetCommentListHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParameter.Wrap(err, "pathparam must be number")
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
