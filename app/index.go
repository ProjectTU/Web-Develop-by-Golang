package app

import (
	"net/http"

	"github.com/jirawat050/GolangWeb/model"
	"github.com/jirawat050/GolangWeb/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	view.Index(w, &view.IndexData{
		List: list,
	})
}
