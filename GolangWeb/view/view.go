package view

import (
	"net/http"
	"github.com/jirawat050/GolangWeb/model"
)
type IndexData struct {
	List []*model.News
}

func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}
func AdminLogin(w http.ResponseWriter, data interface{}){
	render(tpAdminLogin, w ,data)
}
func AdminList(w http.ResponseWriter, data interface{}){
	render(tpAdminList, w ,data)
}
func AdminCreate(w http.ResponseWriter, data interface{}){
	render(tpAdminCreate, w ,data)
}
func AdminEdit(w http.ResponseWriter, data interface{}){
	render(tpAdminEdit, w ,data)
}