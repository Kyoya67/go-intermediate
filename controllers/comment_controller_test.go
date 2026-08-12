package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCommentListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		id         string
		resultCode int
	}{
		{name: "number id", id: "1", resultCode: http.StatusOK},
		{name: "alphabet id", id: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/comment/list/%s", tt.id)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/comment/list/{id}", commentController.GetCommentListHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
